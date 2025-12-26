package config

import "os"

type Config struct {
	Port string

	AuthServiceURL       string
	InventoryServiceURL  string
	PurchasingServiceURL string

	JWTSecret string
}

func Load() *Config {
	return &Config{
		Port:                 getEnv("GATEWAY_PORT", "8080"),
		AuthServiceURL:       getEnv("AUTH_SERVICE_URL", "http://auth-service:8081"),
		InventoryServiceURL:  getEnv("INVENTORY_SERVICE_URL", "http://inventory-service:8082"),
		PurchasingServiceURL: getEnv("PURCHASING_SERVICE_URL", "http://purchasing-service:8083"),
		JWTSecret:            getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
