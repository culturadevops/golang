package routes

import (
	"echo.micro.user/controllers"
	"github.com/labstack/echo"
)

func Opens(e *echo.Echo) {

	e.POST("/login", controllers.Login)

}
