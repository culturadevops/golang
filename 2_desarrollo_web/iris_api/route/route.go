package route

import (
	"cmsx/common"
	"cmsx/controller"
	"cmsx/middleware"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Routes(app *iris.Application) {
	//登录路由
	mvc.New(app.Party("/")).
		Register(commons.SessManager.Start).
		Handle(new(controllers.IndexController))

	//登录路由
	mvc.New(app.Party("/login")).
		Register(commons.SessManager.Start).
		Handle(new(controllers.LoginController))

	//系统路由
	mvc.New(app.Party("/system", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.SystemController))
	//管理员管理
	mvc.New(app.Party("/administrators", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.AdministratorsController))
	//分类管理
	mvc.New(app.Party("/categorys", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.CategorysController))
	//内容管理
	mvc.New(app.Party("/news", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.NewsController))
}
