package database

import (
	"fmt"
	"log"
	"net/url"

	"auth-service/internal/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s@%s:%s/%s?sslmode=disable",
		url.QueryEscape(cfg.DB.User),
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	// log.Println("🔗 FINAL DSN =", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("✅ connected to postgres:", cfg.DB.Name)
	return db, nil
}
