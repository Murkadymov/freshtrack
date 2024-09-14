package handlers

import (
	"encoding/json"
	"fmt"
	"freshtrack/internal/entity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type FreshTrackService interface {
	AddSupply(supply *entity.Supply) error
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

	var supply *entity.Supply

	err := json.NewDecoder(c.Request().Body).Decode(&supply)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			SendError(http.StatusBadRequest, "Bad request", "Invalid JSON format", nil),
		)
	}
	defer c.Request().Body.Close()

	err = h.service.AddSupply(supply)
	if err != nil {
		log.Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			fmt.Sprintf(
				"error: %w",
				SendError(http.StatusInternalServerError, "Internal Server Error", "", nil)),
		)
	}

	err = c.JSON(
		http.StatusCreated,
		fmt.Sprintf("поставка была добавлена %v: ", supply),
	)
	if err != nil {
		log.Error(err)
	}

	return nil
}
