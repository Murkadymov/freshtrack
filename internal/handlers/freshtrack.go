package handlers

import (
	"encoding/json"
	"freshtrack/internal/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FreshTrackService interface {
	AddSupply(supply *entity.Supply) error
	GetSupplyList() ([]entity.Supply, error)
}

type Handler struct {
	service FreshTrackService
}

func NewHandler(service FreshTrackService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddSupply(c echo.Context) error {
	const op = "handlers.freshtrackrepo.AddSupply"

	var supply entity.Supply

	err := json.NewDecoder(c.Request().Body).Decode(&supply)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			h.error(err),
		)
	}
	defer c.Request().Body.Close()

	err = h.service.AddSupply(&supply)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			h.error(err),
		)
	}

	err = c.JSON(
		http.StatusCreated,
		h.ok(echo.Map{"driver": supply.Driver, "goods": supply.Goods, "manufacturer": supply.Manufacturer}),
	)
	if err != nil {
		return echo.NewHTTPError(500, error.Error(err))
	}

	return nil
}

func (h *Handler) GetSupplyList(c echo.Context) error {
	supplyList, err := h.service.GetSupplyList() //get slice of structs
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, error.Error(err))
	}

	return c.JSON(http.StatusOK, h.ok(supplyList))
}
