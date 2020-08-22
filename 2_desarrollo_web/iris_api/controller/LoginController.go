package controllers

import (
	"cmsx/common"
	"cmsx/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html"
	"strings"
	"sync"
)

type LoginController struct {
	Ctx     iris.Context
	Session *sessions.Session
		      sync.Mutex
}

func (c *LoginController) Get() {
	if auth := commons.SessManager.Start(c.Ctx).Get("admin_user"); auth != nil {
		c.Ctx.Redirect("/system/main")
	} else {
		c.Ctx.Redirect("/login/show")
	}
}

func (c *LoginController) GetShow() mvc.View {
	err := c.Ctx.URLParam("err")
	return mvc.View{
		Name:   "user/login.html",
		Layout: "shared/layoutNone.html",
		Data:   iris.Map{"Title": "后台登陆", "err": err},
	}
}

func (c *LoginController) Post() {
	var admin_model model.Admin
	account := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("account")))
	password := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("password")))
	if adminInfo, err := admin_model.AdminLogin(account, password); err == nil { //登录成功
		sessionInfo := make(map[string]interface{})
		sessionInfo["id"] = adminInfo.ID
		sessionInfo["name"] = adminInfo.Account
		c.Session.Set("admin_user", sessionInfo)
		c.Ctx.Redirect("/system/main")
	} else {
		c.Ctx.Redirect("/login/show?err=" + err.Error())
		//登录失败
		//c.Ctx.ViewData("Message", err.Error())
		//c.Ctx.ViewLayout("shared/layoutNone.html")
		//c.Ctx.View(config.GetString("site.DefaultError"))
	}
}

func (c *LoginController) GetLogout() {
	c.Session.Delete("admin_user")
	c.Ctx.Redirect("/login")
}
