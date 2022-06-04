package routers

import (
	"context"
	"database/sql"
	bannerrepository "growdo/src/controller/banner/bannerRepository"
	bannerservices "growdo/src/controller/banner/bannerServices"
	bannerhandler "growdo/src/handler/bannerHandler"
	"growdo/src/helpers/middle"

	"github.com/labstack/echo"
)

func RouteBanner(db *sql.DB, ctx *echo.Echo) {
	c := context.Background()
	repository := bannerrepository.NewRepository(db, c)
	service := bannerservices.NewService(repository)
	handler := bannerhandler.NewHandler(service)

	// api := ctx.Group("/api/v1/user")

	apiAuth := ctx.Group("/api/v1/banner")
	apiAuth.POST("/create", handler.Store, middle.MiddelWareAdmin(2))
	apiAuth.GET("/get", handler.Get)
}
