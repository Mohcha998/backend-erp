package server

import (
	"log"
	"net/http"
	"time"

	_ "auth-service/docs"
	v1 "auth-service/internal/delivery/http/v1"
	"auth-service/internal/infrastructure/config"
	"auth-service/internal/infrastructure/database"
	"auth-service/internal/infrastructure/database/seeders"
	"auth-service/internal/repository"
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// ================= MIGRATION =================
	database.RunMigration(db)

	// ================= SEEDER =================
	if err := seeders.SeedAll(db); err != nil {
		log.Fatal("❌ Seeder failed:", err)
	}

	// ================= REPOSITORIES =================
	userRepo := repository.NewUserRepository(db)
	userRoleRepo := repository.NewUserRoleRepository()
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)

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
		refreshTokenRepo,
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

	// ================= SWAGGER =================
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ================= ROUTES =================
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
		userRepo,
		authUC,
		cfg.JWT.SecretKey,
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
