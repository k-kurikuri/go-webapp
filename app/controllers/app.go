package controllers

import (
	"encoding/json"
	"github.com/k-kurikuri/gogo-done/app/db"
	"github.com/k-kurikuri/gogo-done/app/filters"
	"github.com/k-kurikuri/gogo-done/app/models"
	"github.com/revel/revel"
	"log"
	"strconv"
	"time"
)

type App struct {
	*revel.Controller
}

func init() {
	revel.FilterController(App{}).Insert(filters.AuthFilter, revel.BEFORE, revel.ActionInvoker)
}

func (c App) Index() revel.Result {
	con := db.Connection()

	categories := []models.Category{}

	con.Find(&categories)

	doneLists := []models.DoneList{}
	con.Find(&doneLists)

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
	return c.Render()
}

// 最後にやった事を登録する
func (c App) createDoneList(title string, categoryId uint, postedAt string) error {
	con := db.Connection()
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
