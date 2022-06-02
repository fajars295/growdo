package bannerhandler

import (
	bannerservices "growdo/src/controller/banner/bannerServices"
	bannervalidator "growdo/src/controller/banner/bannerValidator"
	"growdo/src/helpers/componen"
	"growdo/src/helpers/respone"
	"growdo/src/model"

	"github.com/labstack/echo"
)

type handler struct {
	service bannerservices.Service
}

func NewHandler(service bannerservices.Service) *handler {
	return &handler{service: service}
}

func (h *handler) Store(c echo.Context) error {

	var objRequest model.Banner
	if err := c.Bind(&objRequest); err != nil {
		data, status := respone.ERROR(c, "terjadi kesalahan pada system", err)
		return c.JSON(status, data)
	}

	validator := bannervalidator.ValidatorCreate(&objRequest)
	if validator != "" {
		data, status := respone.ERROR(c, validator, nil)
		return c.JSON(status, data)
	}

	save, mes, err := h.service.Create(&objRequest)
	if err != nil {
		data, status := respone.ERROR(c, mes, err)
		return c.JSON(status, data)
	}

	data, status := respone.JSON(c, mes, save)
	return c.JSON(status, data)
}

func (h *handler) Get(c echo.Context) error {

	var con string

	// get status in query parameter
	s := c.QueryParam("status")
	con = "all"

	if s != "" {
		con = "status"
		if s != "true" && s != "false" {
			data, status := respone.ERROR(c, "can only search true or false", nil)
			return c.JSON(status, data)
		}
	}

	cari := &model.FilterCari{
		StatusBanner: componen.StringToBool(s),
	}

	save, mes, err := h.service.All(con, cari)
	if err != nil {
		data, status := respone.ERROR(c, mes, err)
		return c.JSON(status, data)
	}

	data, status := respone.JSON(c, mes, save)
	return c.JSON(status, data)
}
