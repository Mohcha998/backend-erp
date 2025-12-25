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

	// ================= GIN =================
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("❌ trusted proxy error:", err)
	}

	// ================= DATABASE =================
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal("❌ database connection failed:", err)
	}

	// Auto migration
	database.RunMigration(db)

	// ================= REPOSITORIES =================
	userRepo := repository.NewUserRepository(db)
	userRoleRepo := repository.NewUserRoleRepository()

	divisionRepo := repository.NewDivisionRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	menuRepo := repository.NewMenuRepository(db)

	roleMenuRepo := repository.NewRoleMenuRepository(db)
	rolePermissionRepo := repository.NewRolePermissionRepository(db)

	// ================= USECASES =================
	authUC := usecase.NewAuthUsecase(
		db,
		userRepo,
		userRoleRepo,
		cfg.JWT.SecretKey,
	)

	userUC := usecase.NewUserUsecase(db, userRepo)
	divisionUC := usecase.NewDivisionUsecase(divisionRepo)
	roleUC := usecase.NewRoleUsecase(roleRepo)
	menuUC := usecase.NewMenuUsecase(menuRepo)

	roleMenuUC := usecase.NewRoleMenuUsecase(roleMenuRepo)
	rolePermissionUC := usecase.NewPermissionUsecase(rolePermissionRepo)

	// ================= HANDLERS =================
	authHandler := v1.NewAuthHandler(authUC)
	userHandler := v1.NewUserHandler(userUC)
	divisionHandler := v1.NewDivisionHandler(divisionUC)
	roleHandler := v1.NewRoleHandler(roleUC)
	menuHandler := v1.NewMenuHandler(menuUC)

	roleMenuHandler := v1.NewRoleMenuHandler(roleMenuUC)
	rolePermissionHandler := v1.NewRolePermissionHandler(rolePermissionUC)

	// ================= ROUTES =================
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	v1.RegisterRoutes(
		api,
		authHandler,
		userHandler,
		roleHandler,
		menuHandler,
		divisionHandler,
		roleMenuHandler,
		rolePermissionHandler,
	)

	// ================= HTTP SERVER =================
	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("🚀 auth-service running on port", cfg.App.Port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("❌ server error:", err)
	}
}
