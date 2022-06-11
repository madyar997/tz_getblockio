package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tz_getblockio/service"
)

type Handler struct {
	mainNetService *service.MainNetService
}

func NewHandler(netService *service.MainNetService) *Handler{
	return &Handler{mainNetService: netService}
}

func(h Handler) GetMaximumChange(c echo.Context) error {
	address, err := h.mainNetService.GetMaxChange()
	if err !=  nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, address)
}