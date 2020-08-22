package controllers

import (
	"net/http"

	"echo.micro.user/libs"
	"echo.micro.user/middlewares"
	"echo.micro.user/models"
	"github.com/labstack/echo"
)

func Loginx(c echo.Context) error {

	u := new(models.User)

	if err := libs.DecodeValidate(u, c); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
func Login(c echo.Context) error {
	u := new(models.User)
	if err := libs.DecodeValidate(u, c); err != nil {
		return err
	}

	if _, err := u.Login(); err == nil {

		var claims middlewares.JwtCustomClaims
		var data middlewares.CustomData

		t, err := claims.CreateToken(data.NewData(u.Account, true))

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	} else {
		return echo.ErrUnauthorized
		/*return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})*/

	}

}
