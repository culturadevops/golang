package commons

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	config "github.com/spf13/viper"
	"math/rand"
	"strconv"
	"time"
)

/*
session全局共享定义
调用方法
if auth := commons.SessManager.Start(Ctx).Get("admin_user"); auth == nil {
		Ctx.Redirect("/login")
		return
	}
*/
var SessManager = sessions.New(sessions.Config{
	Cookie:  config.GetString("site.SessionCoolieName"),
	Expires: time.Duration(config.GetInt64("SessionExpires")) * time.Hour,
})

/*获取随机整数*/
func GenerateRangeNum(min int, max int) int {
	if min == max {
		return min
	}
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

/*
获取分页html
*/
func GetPageHtml(totalPages, currentPage, total int, path string) string {
	list := ""
	for i := 1; i <= totalPages; i++ {
		active := ""
		if i == currentPage {
			active = "active"
		}
		list += "<li class='paginate_button " + active + "'><a href='" + path + "?page=" + strconv.Itoa(i) + "'>" + strconv.Itoa(i) + "</a></li> "
	}

	previous := ""
	if currentPage == 1 {
		previous = "<li class='paginate_button previous disabled'><a href='#'>Previous</a></li> "
	} else {
		previous = "<li class='paginate_button previous'><a href='" + path + "?page=" + strconv.Itoa(currentPage-1) + "'>Previous</a></li> "
	}

	next := ""
	if currentPage >= totalPages {
		next = "<li class='paginate_button next disabled'><a href='#'>Next</a></li> "
	} else {
		next = "<li class='paginate_button next'><a href='" + path + "?page=" + strconv.Itoa(currentPage+1) + "'>Next</a></li> "
	}

	html := "<div class='box-footer clearfix'><div class='dataTables_info'>总共" + strconv.Itoa(total) + "条记录/" + strconv.Itoa(totalPages) + "页</div>" +
		"<ul class='pagination pagination-sm no-margin pull-right'> " +
		previous +
		list +
		next +
		"</ul> " +
		"</div>"
	return html
}

/*
错误页面显示
*/
func DefaultErrorShow(msg string, Ctx iris.Context) {
	config.SetConfigName("app")
	Ctx.ViewData("Code", strconv.Itoa(Ctx.GetStatusCode()))
	Ctx.ViewData("Title", "信息异常")
	Ctx.ViewData("Message", msg)
	Ctx.View(config.GetString("site.DefaultError"))
}

/*
MVC错误页面显示
*/
func MvcError(msg string, Ctx iris.Context) mvc.View {
	config.SetConfigName("app")
	return mvc.View{
		Name: config.GetString("site.DefaultError"),
		Data: iris.Map{
			"Title":   "信息异常",
			"Code":    strconv.Itoa(Ctx.GetStatusCode()),
			"Message": msg,
		},
	}
}
