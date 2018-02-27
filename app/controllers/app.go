package controllers

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/k-kurikuri/gogo-done/app/db"
	"github.com/k-kurikuri/gogo-done/app/filters"
	"github.com/k-kurikuri/gogo-done/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func init() {
	revel.FilterController(App{}).Insert(filters.AuthFilter, revel.BEFORE, revel.ActionInvoker)
}

func (c App) Index() revel.Result {
	con := db.Connection()
	defer con.Close()

	categories := []models.Category{}
	con.Find(&categories)

	doneLists := []models.DoneList{}
	con.Find(&doneLists)

	ids := []uint{}
	for _, doneList := range doneLists {
		ids = append(ids, doneList.Id)
	}

	rows, _ := con.Table("done_list_histories").
		Select("done_list_id, MAX(posted_at)").
		Where("done_list_id IN (?)", ids).
		Group("done_list_id").
		Rows()

	defer rows.Close()

	for rows.Next() {
		var doneListId uint
		var latestPost time.Time
		rows.Scan(&doneListId, &latestPost)
		// todo: viewModelを生成する
	}

	return c.Render(categories, doneLists)
}

func (c App) Create() revel.Result {
	title := c.Params.Form.Get("title")
	categoryIdStr := c.Params.Form.Get("category")
	postedAt := c.Params.Form.Get("postedAt")

	// TODO: validate

	categoryId, _ := strconv.Atoi(categoryIdStr)
	// 配列indexをincrement
	categoryId++

	if err := c.createDoneList(title, uint(categoryId), postedAt); err != nil {
		log.Panic("create donelist failed")
	}

	return c.Redirect(App.Index)
}

func (c App) Detail() revel.Result {
	doneListId := c.Params.Route.Get("id")

	con := db.Connection()
	defer con.Close()

	doneList := models.DoneList{}
	con.First(&doneList, doneListId)

	doneListHistories := []models.DoneListHistory{}
	con.Where("done_list_id = ?", doneListId).Find(&doneListHistories)

	return c.Render(doneList, doneListHistories)
}

func (c App) Update() revel.Result {
	// TODO: auth user's record validate
	historyId := c.Params.Form.Get("pk")
	note := c.Params.Form.Get("value")

	con := db.Connection()
	defer con.Close()

	doneHistory := models.DoneListHistory{}
	con.Model(&doneHistory).Where("id = ?", historyId).Update("note", note)

	return c.RenderJSON("success")
}

// 最後にやった事を登録する
func (c App) createDoneList(title string, categoryId uint, postedAt string) error {
	con := db.Connection()
	defer con.Close()

	tran := con.Begin()

	user, _ := c.sessionUser()

	doneList := models.DoneList{
		Title:      title,
		CategoryId: categoryId,
		UserId:     user.Id,
	}

	if err := tran.Create(&doneList).Error; err != nil {
		tran.Rollback()
		return err
	}

	t, _ := time.Parse("2006-01-02", postedAt)
	log.Println(t)
	doneListHistory := models.DoneListHistory{
		DoneListId: doneList.Id,
		Note:       "イベント作成済み",
		PostedAt:   t,
	}

	if err := tran.Create(&doneListHistory).Error; err != nil {
		tran.Rollback()
		return err
	}

	tran.Commit()

	return nil
}

// TODO: controller依存せず使用したい
func (c App) sessionUser() (models.User, error) {
	jsonStr := c.Session["user"]

	jsonBytes := ([]byte)(jsonStr)

	user := models.User{}

	err := json.Unmarshal(jsonBytes, &user)

	return user, err
}
