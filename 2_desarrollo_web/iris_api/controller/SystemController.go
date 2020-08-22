package controllers

import (
	"cmsx/common"
	"cmsx/libs"
	"cmsx/model"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html"
	"runtime"
	"strings"
)

var admin_model model.Admin

type SystemController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *SystemController) GetMain() mvc.View {
	return mvc.View{
		Name: "system/main.html",
		Data: iris.Map{
			"Title":     "系统概况",
			"CpuNum":    runtime.NumCPU(),
			"GoVersion": runtime.Version(),
			"Goos":      runtime.GOOS,
			"GoRoot":    runtime.GOROOT(),
		},
	}
}

func (c *SystemController) GetUpdatePassword() mvc.View {
	return mvc.View{
		Name: "system/updatePassword.html",
		Data: iris.Map{
			"Title": "密码修改",
		},
	}
}

func (c *SystemController) PostUpdatePassword() {
	password := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("password")))
	Repassword := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("Repassword")))
	admin_user, _ := c.Session.Get("admin_user").(map[string]interface{})
	admin_id, _ := admin_user["id"].(uint)
	if err := admin_model.AdminPasswodUpdate(admin_id, password, Repassword); err == nil {
		c.Ctx.Redirect("/system/main")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *SystemController) GetUpdateAdmin() mvc.View {
	return mvc.View{
		Name: "system/updateAdmin.html",
		Data: iris.Map{
			"Title": "资料修改",
		},
	}
}

func (c *SystemController) PostUpdateAdmin() {
	admin_user, _ := c.Session.Get("admin_user").(map[string]interface{})
	admin_id, _ := admin_user["id"].(uint)
	err, filePath := libs.UploadFile("headico", c.Ctx)
	if err == false {
		commons.DefaultErrorShow(filePath, c.Ctx)
		return
	}
	if err := admin_model.AdminUpdate(c.Ctx.FormValues(), admin_id, filePath); err == nil {
		c.Ctx.Redirect("/system/update/admin")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
