package controllers

import (
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
	return c.Render()
}

func (c App) Create() revel.Result {
	title := c.Params.Form.Get("title")
	categoryIdStr := c.Params.Form.Get("category")
	postedAt := c.Params.Form.Get("postedAt")

	// TODO: validate

	categoryId, _ := strconv.Atoi(categoryIdStr)

	if err := c.createDoneList(title, uint(categoryId), postedAt); err != nil {
		log.Panic("create donelist failed")
	}

	res := make(map[string]interface{})
	res["isResult"] = true

	return c.RenderJSON(res)
}

// 最後にやった事を登録する
func (c App) createDoneList(title string, categoryId uint, postedAt string) error {
	con := db.Connection()
	tran := con.Begin()

	doneList := models.DoneList{
		Title:      title,
		CategoryId: categoryId,
		//UserId: TODO
	}

	if err := tran.Create(&doneList).Error; err != nil {
		tran.Rollback()
		return err
	}

	t, _ := time.Parse("2006-01-02", postedAt)
	log.Println(t)
	doneListHistory := models.DoneListHistory{
		DoneListId: doneList.Id,
		PostedAt:   t,
	}

	if err := tran.Create(&doneListHistory).Error; err != nil {
		tran.Rollback()
		return err
	}

	tran.Commit()

	return nil
}
