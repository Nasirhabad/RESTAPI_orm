package middlewares

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey string
}

func (j *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(j.SecretKey),
	}

}

func GetUser(c echo.Context) (*JWTCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return nil, errors.New("invalid Token")
	}

	claims := user.Claims.(*JWTCustomClaims)
	return claims, nil
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userdata, err := GetUser(c)
		isInvalid := userdata == nil || err != nil
		if isInvalid {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "invalid Token",
			})
		}
		return next(c)
	}

}
