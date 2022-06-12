package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tz_getblockio/internal/model"
	"tz_getblockio/internal/service"
)

type Handler struct {
	mainNetService *service.MainNetService
}

func NewHandler(netService *service.MainNetService) *Handler {
	return &Handler{mainNetService: netService}
}

// GetMaximumChange  godoc
// @Description  get the address of the account which balance changed the most(also provides the receiver address) over the last 100 blocks
// @Tags         api/v1
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Result
// @Failure      400  {object}  model.Error
// @Router       /maximum-change [get]
func (h Handler) GetMaximumChange(c echo.Context) error {
	address, err := h.mainNetService.GetMaxChange()
	if err != nil {
		errMessage := model.Error{Message: err.Error()}
		return c.JSON(http.StatusInternalServerError, errMessage)
	}
	return c.JSON(http.StatusOK, address)
}
