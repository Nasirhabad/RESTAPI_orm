package controllers

import (
	"my-orm/models"
	"my-orm/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PackageController struct {
	service services.PackageService
}

func InitPackageController() PackageController {
	return PackageController{
		service: services.PackageService{},
	}
}
func (pc *PackageController) GetAll(c echo.Context) error {
	pkg, err := pc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error occured whe fetching data",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": pkg,
	})
}

func (pc *PackageController) GetByID(c echo.Context) error {
	id := c.Param("id")
	pkg, err := pc.service.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error occured whe fetching data",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": pkg,
	})
}

func (pc *PackageController) Create(c echo.Context) error {
	var packageReq models.PackageRequest
	if err := c.Bind(&packageReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to add package",
		})
	}

	pkg, err := pc.service.Create(packageReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error occured whe creating package data",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Created successfuly",
		"package": pkg,
	})
}

func (pc *PackageController) Update(c echo.Context) error {

	id := c.Param("id")
	var pkgReq models.PackageRequest
	if err := c.Bind(&pkgReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "update package failed",
		})
	}
	pkg, err := pc.service.Update(id, pkgReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "update failed",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "package updated",
		"package": pkg,
	})

}

func (pc *PackageController) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := pc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "delete failed",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "delete successful",
	})

}
