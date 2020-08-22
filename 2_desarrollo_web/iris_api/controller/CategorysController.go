package controllers

import (
	"cmsx/common"
	"cmsx/model"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

var CategoryModel = model.Category{}

type CategorysController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *CategorysController) Get() mvc.View {
	Category := model.Category{}
	list := Category.List()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)
	return mvc.View{
		Name: "category/list.html",
		Data: iris.Map{
			"Title": "分类列表",
			"list":  list,
		},
	}
}

func (c *CategorysController) GetAddCategory() mvc.View {
	Category := model.Category{}
	list := Category.List()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)
	return mvc.View{
		Name: "category/addCategory.html",
		Data: iris.Map{
			"Title": "新增分类",
			"list":  list,
		},
	}
}

func (c *CategorysController) PostAddCategory() {
	if err := CategoryModel.CategoryAdd(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/categorys")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *CategorysController) GetUpdateCategoryBy(id uint) mvc.View {
	categoryInfo, err := CategoryModel.CategoryInfo(id)
	if err != nil {
		return commons.MvcError(err.Error(), c.Ctx)
	}
	Category := model.Category{}
	list := Category.List()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)

	return mvc.View{
		Name: "category/updateCategory.html",
		Data: iris.Map{
			"Title":              "分类修改",
			"UpdateCategoryInfo": categoryInfo,
			"list":               list,
		},
	}
}

func (c *CategorysController) PostUpdateCategory() {
	if err := CategoryModel.CategoryUpdate(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/categorys")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *CategorysController) GetDelCategoryBy(id uint) {
	if err := CategoryModel.CategoryDel(id); err == nil {
		c.Ctx.Redirect("/categorys")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
