package main

import (
	"encoding/json"
	"microconsultor/libs"
	"microconsultor/routes"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	config "github.com/spf13/viper"
)

func init() {
	config.AddConfigPath("./configs")
	config.SetConfigName("myProp")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	formData := url.Values{
		"user":     {config.GetString("default.user")},
		"password": {config.GetString("default.passwd")},
	}
	formData = url.Values{
		"Micro":   {config.GetString("default.name")},
		"Env":     {config.GetString("default.env")},
		"Version": {config.GetString("default.version")},
	}
	var authAddr string = config.GetString("default.authaddr") + "/api/v1/get"
	resp, _ := http.PostForm(authAddr, formData)
	defer resp.Body.Close()
	//data, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(data))
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	dbConfig := libs.DbConfig{
		string(result["host"]),
		string(result["port"]),
		string(result["database"]),
		string(result["user"]),
		string(result["password"]),
		string(result["charset"]),
		10,
		100,
		//result["MaxIdleConns"),
		//result["MaxOpenConns"),
	}
	libs.DB = dbConfig.InitDB()
	if config.GetBool("default.sql_log") {
		libs.DB.LogMode(true)
	}

}

func main() {

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//log.Debug("err", err)
	//log.Debug("resp", resp.Body)
	//Routers
	routes.Opens(e)

	e.Logger.Fatal(e.Start(":8881"))
}
