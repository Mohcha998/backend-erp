package database

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"auth-service/internal/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

// ==============================
// INIT CONNECTION
// ==============================
func NewPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		url.QueryEscape(cfg.DB.User),
		url.QueryEscape(cfg.DB.Password),
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	var db *gorm.DB
	var err error

	for i := 1; i <= 15; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Connected to PostgreSQL")
			SetDB(db)
			setPool(db)
			return db, nil
		}

		log.Printf("⏳ Waiting for DB... (%d/15)\n", i)
		time.Sleep(2 * time.Second)
	}

	return nil, err
}

// ==============================
// GLOBAL DB ACCESS
// ==============================
func SetDB(db *gorm.DB) {
	dbInstance = db
}

func GetDB() *gorm.DB {
	return dbInstance
}

// ==============================
// CONNECTION POOL
// ==============================
func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("⚠️ Failed to get DB instance:", err)
		return
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
}
