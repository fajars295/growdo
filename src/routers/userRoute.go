package routers

import (
	"database/sql"
	"fmt"
	"growdo/src/controller/user/userRepository"
	"growdo/src/controller/user/userServices"
	userhandler "growdo/src/handler/userHandler"
	"growdo/src/helpers/middle"

	"github.com/labstack/echo"
)

func RouteUser(db *sql.DB, ctx *echo.Echo) {

	fmt.Println("mauk route")

	repository := userRepository.NewRepository(db)
	service := userServices.NewService(repository)
	handler := userhandler.NewHandler(service)

	api := ctx.Group("/api/v1/user")
	api.POST("/register", handler.Store)
	api.POST("/login", handler.Login)

	apiAuth := ctx.Group("/api/v1/user")
	apiAuth.GET("/profile", handler.Profile, middle.IsLogin)
}
