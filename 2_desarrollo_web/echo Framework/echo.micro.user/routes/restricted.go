package routes

import (
	"net/http"

	"echo.micro.user/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ControllerRestricted(c echo.Context) error {
	//user := c.Get("user").(*jwt.Token)
	//claims := user.Claims.(*middlewares.JwtCustomClaims)
	//name := claims.Name
	return c.String(http.StatusOK, "Welcome !")
}
func Restricted(e *echo.Echo) {

	jwtconfig := middlewares.JwtInitConfig()
	r := e.Group("/api/v1/")
	r.Use(middleware.JWTWithConfig(jwtconfig))
	//r.GET("", ControllerRestricted)
}
