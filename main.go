package main

import (
	"freshtrack/internal/config"
	"freshtrack/internal/handlers"
	"freshtrack/internal/repository/freshtrackrepo"
	"freshtrack/internal/repository/postgre"
	"freshtrack/internal/service/freshtrackservice"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.MustLoad()

	db := postgre.ConnectToDB(cfg)

	repo := freshtrackrepo.NewFreshTrackRepository(db)
	service := freshtrackservice.NewService(repo)
	handler := handlers.NewHandler(service)

	//logg := logger2.NewLogger()

	e := echo.New()

	e.POST("/supply", handler.AddSupply)

	e.Start(":8080")

}
