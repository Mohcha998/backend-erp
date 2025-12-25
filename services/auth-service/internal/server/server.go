package server

import (
	"log"
	"net/http"
	"time"

	v1 "auth-service/internal/delivery/http/v1"
	"auth-service/internal/infrastructure/config"
	"auth-service/internal/infrastructure/database"
	"auth-service/internal/repository"
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Run() {
	// ================= CONFIG =================
	cfg := config.Load()

	// ================= GIN ENGINE =================
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})

	// ================= DATABASE =================
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal("❌ failed to connect database:", err)
	}

	// ================= REPOSITORIES =================
	userRepo := repository.NewUserRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db) // wajib untuk auth usecase

	// ================= USECASE =================
	authUC := usecase.NewAuthUsecase(db, userRepo, userRoleRepo, cfg.JWT.SecretKey)

	// ================= HANDLER =================
	authHandler := v1.NewAuthHandler(authUC)

	// ================= HEALTH CHECK =================
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// ================= ROUTES =================
	api := r.Group("/api")
	v1.RegisterAuthRoutes(api, authHandler)

	// ================= SERVER =================
	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("🌐 auth-service running on port", cfg.App.Port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("❌ server error:", err)
	}
}
