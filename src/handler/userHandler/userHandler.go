package userhandler

import (
	"growdo/src/controller/user/userServices"
	"growdo/src/controller/user/userValidator"
	"growdo/src/helpers/componen"
	"growdo/src/helpers/middle"
	"growdo/src/helpers/respone"
	"growdo/src/model"

	"github.com/labstack/echo"
)

type handler struct {
	service userServices.Service
}

func NewHandler(service userServices.Service) *handler {
	return &handler{service: service}
}

func (h *handler) Store(c echo.Context) error {

	var objRequest model.Users
	if err := c.Bind(&objRequest); err != nil {
		data, status := respone.ERROR(c, "terjadi kesalahan pada system", err)
		return c.JSON(status, data)
	}

	validator := userValidator.ValidatorCreate(&objRequest)
	if validator != "" {
		data, status := respone.ERROR(c, validator, nil)
		return c.JSON(status, data)
	}

	objRequest.Roles = componen.CheckIfRoles(objRequest.Roles)

	save, mes, err := h.service.Create(&objRequest)
	if err != nil {
		data, status := respone.ERROR(c, mes, err)
		return c.JSON(status, data)
	}

	data, status := respone.JSON(c, mes, save)
	return c.JSON(status, data)
}

func (h *handler) Login(c echo.Context) error {

	var objRequest model.Login
	if err := c.Bind(&objRequest); err != nil {
		data, status := respone.ERROR(c, "terjadi kesalahan pada system", err)
		return c.JSON(status, data)
	}

	validator := userValidator.ValidatorLogin(&objRequest)
	if validator != "" {
		data, status := respone.ERROR(c, validator, nil)
		return c.JSON(status, data)
	}

	save, mes, err := h.service.Login(&objRequest)
	if err != nil {
		data, status := respone.ERROR(c, mes, err)
		return c.JSON(status, data)
	}

	message := map[string]interface{}{
		"token": middle.GenerateToken(int(save.Id), save),
		"user":  save,
	}

	data, status := respone.JSON(c, mes, message)
	return c.JSON(status, data)
}

func (h *handler) Profile(c echo.Context) error {

	save, mes, err := h.service.Profile(c.Get("id").(int))
	if err != nil {
		data, status := respone.ERROR(c, mes, err)
		return c.JSON(status, data)
	}

	data, status := respone.JSON(c, mes, save)
	return c.JSON(status, data)
}
