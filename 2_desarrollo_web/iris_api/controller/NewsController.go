package controllers

import (
	"cmsx/common"
	"cmsx/model"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strconv"
	"strings"
)

type NewsController struct {
	Ctx     iris.Context
	Session *sessions.Session
	News    model.News
}

func (c *NewsController) Get() mvc.View {
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages := c.News.List(page)
	Category := model.Category{}
	for k, v := range list {
		CategoryName := ""
		if val, err := Category.CategoryMoreInfo(v.Category_id); err == nil {
			for _, vv := range val {
				CategoryName += vv.Name + ","
			}
		}
		list[k].CategoryName = strings.TrimRight(CategoryName, ",")
	}
	return mvc.View{
		Name: "news/list.html",
		Data: iris.Map{
			"Title":    "内容列表",
			"list":     list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}
}

func (c *NewsController) GetAddNews() mvc.View {
	Category := model.Category{}
	list := Category.List()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)
	return mvc.View{
		Name: "news/addNews.html",
		Data: iris.Map{
			"Title": "新增内容",
			"list":  list,
		},
	}
}

func (c *NewsController) PostAddNews() {
	if err := c.News.NewsAdd(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/news")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *NewsController) GetUpdateNewsBy(id uint) mvc.View {
	NewsInfo, err := c.News.NewsInfo(id)
	if err != nil {
		return commons.MvcError(err.Error(), c.Ctx)
	}
	Category := model.Category{}
	list := Category.List()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)

	CategoryIds := []int{}
	for _, v := range strings.Split(NewsInfo.Category_id, ",") {
		_v, _ := strconv.Atoi(v)
		CategoryIds = append(CategoryIds, _v)
	}

	return mvc.View{
		Name: "news/updateNews.html",
		Data: iris.Map{
			"Title":          "内容修改",
			"UpdateNewsInfo": NewsInfo,
			"CategoryIds":    CategoryIds,
			"list":           list,
		},
	}
}

func (c *NewsController) PostUpdateNews() {
	if err := c.News.NewsUpdate(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/news")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *NewsController) GetDelNewsBy(id uint) {
	if err := c.News.NewsDel(id); err == nil {
		c.Ctx.Redirect("/news")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
