package middlewares

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// jwtCustomClaims are custom claims extending default ones.
// del futuro para cuando sepa usar interfaces has este objeto algo mas interesante

type CustomData struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

type JwtCustomClaims struct {
	Data CustomData
	jwt.StandardClaims
}

var signingKey = []byte("My Secret")

func (t *CustomData) NewData(username string, admin bool) CustomData {
	return CustomData{
		Name:  username,
		Admin: admin,
	}
}

func JwtInitConfig() middleware.JWTConfig {
	// Configure middleware with the custom claims type
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: signingKey,
	}

}
func (this *JwtCustomClaims) CreateToken(Data CustomData) (string, error) {
	this = &JwtCustomClaims{
		Data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	return this.signedClaim()
}

func (this *JwtCustomClaims) signedClaim() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, this)
	return token.SignedString(signingKey)

}

//example for login controller use in a controller for validate credencial
func LoginController(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}
	var claims JwtCustomClaims
	var data CustomData

	t, err := claims.CreateToken(data.NewData(username, true))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
