package postgre

import (
	"database/sql"
	"fmt"
	"freshtrack/internal/config"
	"log"
)

func ConnectToDB(cfg *config.Config) *sql.DB {

	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DB,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	))
	if err != nil {
		log.Fatalf("error openning db: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error verifying conenctiong to db: %s", err)
	}

	return db
}
