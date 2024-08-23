package main

import (
	"my-orm/controllers"
	"my-orm/database"
	"my-orm/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

func main() {

	database.InitDB()

	database.MigrateDB()

	e := echo.New()

	loggerConfig := middlewares.LoggerConfig{

		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()
	e.Use(loggerMiddleware)

	jwtConfig := middlewares.JWTConfig{
		SecretKey: "secret",
	}

	authmiddlewareConfig := jwtConfig.Init()
	userController := controllers.InitUserController()

	packageController := controllers.InitPackageController()
	e.POST("/api/v1/register", userController.Register)
	e.POST("/api/v1/login", userController.Login)

	packageRoutes := e.Group("/api/v1", echojwt.WithConfig(authmiddlewareConfig), middlewares.VerifyToken)

	packageRoutes.GET("/packages", packageController.GetAll)
	packageRoutes.GET("/packages/:id", packageController.GetByID)
	packageRoutes.POST("/packages", packageController.Create)
	packageRoutes.PUT("/packages/:id", packageController.Update)
	packageRoutes.DELETE("/packages/:id", packageController.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
