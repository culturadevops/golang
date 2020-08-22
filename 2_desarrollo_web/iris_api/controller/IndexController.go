package controllers

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"

	"github.com/kataras/iris/sessions"
)

type IndexController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *IndexController) Get()  {
	 c.Ctx.JSON(iris.Map{
		"status": 200,
	"Title":           "资料修改",
		
	})
	
/*	return mvc.View{
		Name:   "index/index.html",
		Layout: "shared/layoutFront.html",
		Data:   iris.Map{"Title": "index"},
	}*/
}
