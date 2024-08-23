package controllers

import (
	"my-orm/models"
	"my-orm/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserController() UserController {
	return UserController{
		service: services.UserService{},
	}
}

func (u *UserController) Register(c echo.Context) error {
	var userReq models.UserRequest
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "register failed",
		})
	}

	user, err := u.service.Register(userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "register failed",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "register success",
		"user":    user,
	})
}

func (u *UserController) Login(c echo.Context) error {
	var userReq models.UserRequest
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "login failed",
		})
	}

	token, err := u.service.Login(userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "login failed",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "login success",
		"user":    token,
	})
}
