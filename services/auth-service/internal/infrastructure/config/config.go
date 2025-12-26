package config

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App      AppConfig
	DB       DBConfig
	JWT      JWTConfig
	Security SecurityConfig
}

type AppConfig struct {
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
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

type SecurityConfig struct {
	AccessTokenExp  string `yaml:"access_token_exp"`
	RefreshTokenExp string `yaml:"refresh_token_exp"`
}

func Load() *Config {
<<<<<<< HEAD
	return &Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "8081"),
		},
		// DB: DBConfig{
		// 	Host:     getEnv("PGHOST", "127.0.0.1"),
		// 	Port:     getEnv("PGPORT", "5432"),
		// 	User:     getEnv("PGUSER", "postgres"),
		// 	Password: getEnv("PGPASSWORD", ""),
		// 	Name:     getEnv("PGDATABASE", "erp_db"),
		// },
		DB: DBConfig{
			Host:     getEnv("PGHOST", "localhost"),
			Port:     getEnv("PGPORT", "5432"),
			User:     getEnv("PGUSER", "postgres"),
			Password: getEnv("PGPASSWORD", "123"),
			Name:     getEnv("PGDATABASE", "erp_db"),
		},
		JWT: JWTConfig{
			SecretKey: getEnv("JWT_SECRET", "secret"),
		},
=======
	_ = godotenv.Load()

	cfg := &Config{}

	// Load yaml
	file, _ := os.ReadFile("config/config.yaml")
	_ = yaml.Unmarshal(file, cfg)

	// Override from ENV
	cfg.App.Port = getEnv("APP_PORT", cfg.App.Port)

	cfg.DB = DBConfig{
		Host:     getEnv("PGHOST", "127.0.0.1"),
		Port:     getEnv("PGPORT", "5432"),
		User:     getEnv("PGUSER", "MRCorp"),
		Password: getEnv("PGPASSWORD", ""),
		Name:     getEnv("PGDATABASE", "erp_db"),
>>>>>>> 8ef03f98876a76569ae9f44783c4e6430826df7d
	}

	cfg.JWT.SecretKey = getEnv("JWT_SECRET", "secret")

	return cfg
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
