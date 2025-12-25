package config

import (
	"os"
)

type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
}

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	SecretKey string
}

func Load() *Config {
	cfg := &Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "8081"),
		},
		DB: DBConfig{
			Host:     getEnv("PGHOST", "127.0.0.1"),
			Port:     getEnv("PGPORT", "5432"),
			User:     getEnv("PGUSER", "MRCorp"),
			Password: os.Getenv("PGPASSWORD"),
			Name:     getEnv("PGDATABASE", "erp_db"),
		},
		JWT: JWTConfig{
			SecretKey: getEnv("JWT_SECRET", "secret"),
		},
	}

	// log.Println("✅ config loaded")
	// log.Println("📦 DB =", cfg.DB.Name)

	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
