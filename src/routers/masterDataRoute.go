package routers

import (
	"context"
	"database/sql"
	masterDatarepository "growdo/src/controller/masterData/masterDataRepository"
	"growdo/src/controller/masterData/masterDataServices"
	"growdo/src/handler/masterDataHandler"
	"growdo/src/helpers/middle"

	"github.com/labstack/echo"
)

func RouteMasterData(db *sql.DB, ctx *echo.Echo) {
	c := context.Background()
	repository := masterDatarepository.NewRepository(db, c)
	service := masterDataServices.NewService(repository)
	handler := masterDataHandler.NewHandler(service)

	apiAuth := ctx.Group("/api/v1/master-data")
	apiAuth.POST("/create", handler.Store, middle.MiddelWareAdmin(2))
	apiAuth.GET("/get/:kolom", handler.FindOne)
}
