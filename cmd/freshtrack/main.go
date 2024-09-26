package main

import (
	"freshtrack/internal/config"
	"freshtrack/internal/http-server/handlers/filemanager"
	"freshtrack/internal/http-server/handlers/freshtrack"
	logger2 "freshtrack/internal/http-server/middleware/logger"
	"freshtrack/internal/pkg/logger"
	"freshtrack/internal/repository/freshtrackrepo"
	"freshtrack/internal/repository/postgre"
	"freshtrack/internal/service/freshtrackservice"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	cfg := config.MustLoad()

	db := postgre.ConnectToDB(cfg)

	res, err := db.Exec(`
		BEGIN TRANSACTION;

		CREATE TABLE IF NOT EXISTS drivers(
		driver_id SERIAL PRIMARY KEY,
		tractor_number VARCHAR(30) NOT NULL,
		trail_number VARCHAR(30) NOT NULL,
		driver_number VARCHAR(30) NOT NULL);
		
		CREATE TABLE IF NOT EXISTS goods(
		goods_id SERIAL PRIMARY KEY,
		cargo VARCHAR(30) NOT NULL);
		
		CREATE TABLE IF NOT EXISTS manufacturer(
		manufacturer_id SERIAL PRIMARY KEY,
		name VARCHAR(30) NOT NULL,
		origin VARCHAR(30) NOT NULL);
		    
		CREATE TABLE IF NOT EXISTS supplies(
		    supply_id SERIAL PRIMARY KEY,
		    driver_id INT NOT NULL,
		    goods_id INT NOT NULL,
		    manufacturer_id INT NOT NULL,
		    FOREIGN KEY (driver_id) REFERENCES drivers(driver_id),
		    FOREIGN KEY (goods_id) REFERENCES goods(goods_id),
		    FOREIGN KEY (manufacturer_id) REFERENCES manufacturer(manufacturer_id));
		
		COMMIT TRANSACTION;`)
	if err != nil {
		log.Fatalf("error executing query: %v\n result: %v", err, res)
	}
	log.Println("Result:", res)

	repo := freshtrackrepo.NewFreshTrackRepository(db)
	service := freshtrackservice.NewService(repo)
	handler := freshtrack.NewHandler(service)
	fileManager := filemanager.NewFileManager()

	log := logger.NewLogger()

	e := echo.New()

	e.Static("/", "E:\\Projects\\freshtrack\\public")

	e.POST("/supply", logger2.NewMiddlewareLogger(log, handler.AddSupply))
	e.GET("/supply", handler.GetSupplyList)
	e.POST("/upload", fileManager.UploadFile(log))

	if err := e.Start(":8080"); err != nil {
		panic("failed to start a server")
	}

}
