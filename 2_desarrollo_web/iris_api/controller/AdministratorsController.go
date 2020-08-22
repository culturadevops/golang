package controllers

import (
	"cmsx/common"
	"cmsx/libs"
	"cmsx/model"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html"
	"strconv"
	"strings"
)

type AdministratorsController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *AdministratorsController) Get() mvc.View {
	Admin := model.Admin{}
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages := Admin.List(page)
	return mvc.View{
		Name: "administrators/list.html",
		Data: iris.Map{
			"Title":    "管理员列表",
			"list":     list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}
}

func (c *AdministratorsController) GetUpdateAdminBy(id uint) mvc.View {
	adminInfo, err := admin_model.AdminInfo(id)
	if err != nil {
		return commons.MvcError(err.Error(), c.Ctx)
	}

	if adminInfo.Headico == "" {
		adminInfo.Headico = "/public/adminlit/dist/img/user2-160x160.jpg"
	}
	return mvc.View{
		Name: "administrators/updateAdmin.html",
		Data: iris.Map{
			"Title":           "资料修改",
			"UpdateAdminInfo": adminInfo,
		},
	}
}

func (c *AdministratorsController) PostUpdateAdmin() {
	err, filePath := libs.UploadFile("headico", c.Ctx)
	if err == false {
		commons.DefaultErrorShow(filePath, c.Ctx)
		return
	}
	var postValues map[string][]string
	postValues = c.Ctx.FormValues()
	admin_id := postValues["id"][0]
	int_admin_id, _ := strconv.Atoi(admin_id)
	delete(postValues, "id")
	if err := admin_model.AdminUpdate(postValues, uint(int_admin_id), filePath); err == nil {
		c.Ctx.Redirect("/administrators/update/admin/" + admin_id)
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *AdministratorsController) GetAddAdmin() mvc.View {
	return mvc.View{
		Name: "administrators/addAdmin.html",
		Data: iris.Map{
			"Title": "新增管理员",
		},
	}
}

func (c *AdministratorsController) PostAddAdmin() {
	err, filePath := libs.UploadFile("headico", c.Ctx)
	if err == false {
		commons.DefaultErrorShow(filePath, c.Ctx)
		return
	}

	if err := admin_model.AddUpdate(c.Ctx.FormValues(), filePath); err == nil {
		c.Ctx.Redirect("/administrators")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *AdministratorsController) GetUpdatePasswordBy(id uint) mvc.View {
	adminInfo, err := admin_model.AdminInfo(id)
	if err != nil {
		return commons.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "administrators/updatePassword.html",
		Data: iris.Map{
			"Title":   "密码修改",
			"Id":      id,
			"Account": adminInfo.Account,
		},
	}
}

func (c *AdministratorsController) PostUpdatePassword() {
	id := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("id")))
	password := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("password")))
	Repassword := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("Repassword")))
	int_admin_id, _ := strconv.Atoi(id)
	if err := admin_model.AdminPasswodUpdate(uint(int_admin_id), password, Repassword); err == nil {
		c.Ctx.Redirect("/administrators")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *AdministratorsController) GetDelAdminBy(id uint) {
	if err := admin_model.AdminDel(id); err == nil {
		c.Ctx.Redirect("/administrators")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
