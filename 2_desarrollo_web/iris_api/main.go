package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	config "github.com/spf13/viper"

	"cmsx/libs"

	"cmsx/route"
	"log"
	"strconv"

)

func init() {
	config.AddConfigPath("./configs")
	config.SetConfigName("mysql")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbConfig := libs.DbConfig{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
		config.GetString("default.charset"),
		config.GetInt("default.MaxIdleConns"),
		config.GetInt("default.MaxOpenConns"),
	}
	libs.DB = dbConfig.InitDB()
	if config.GetBool("default.sql_log") {
		libs.DB.LogMode(true)
	}
}

func main() {
	app := iris.New()
	config.SetConfigName("app")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误, %s", err)
	}

	app.Favicon("./favicon.ico")
	app.Use(iris.Gzip)

	//（可选）添加两个内置处理程序
	//可以从任何与http相关的panics中恢复
	//并将请求记录到终端。

// (opcional) agrega dos controladores integrados
// Puede recuperarse de cualquier pánico relacionado con http
// Grabe la solicitud en la terminal.
	app.Use(recover.New())
	app.Use(logger.New())


	//设置错误模版
	app.OnAnyErrorCode(func(ctx iris.Context) {
		_, err := ctx.HTML("<center>Lo siento! Error de página actual, código de error:" + strconv.Itoa(ctx.GetStatusCode()) + "</center>")
		if err != nil {
			log.Fatalf("			Error interno %s", err)
		}
	})

	route.Routes(app)

	//应用配置文件
	app.Configure(iris.WithConfiguration(iris.YAML("./configs/iris.yml")))

	//Run
	www := app.Party("www.")
	{
		currentRoutes := app.GetRoutes()
		for _, r := range currentRoutes {
			www.Handle(r.Method, r.Tmpl().Src, r.Handlers...)
		}
	}
	err := app.Run(iris.Addr(config.GetString("server.domain") + ":" + config.GetString("server.port")))
	if err != nil {
		log.Fatalf("El servicio no pudo iniciarse, código de error %s", err)
	}
}
