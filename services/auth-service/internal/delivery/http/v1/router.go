package v1

import (
	"auth-service/internal/middleware"
	"auth-service/internal/repository"
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes godoc
// @Summary Register all API routes
// @Description Main router for Auth Service
func RegisterRoutes(
	rg *gin.RouterGroup,

	// handlers
	auth *AuthHandler,
	user *UserHandler,
	role *RoleHandler,
	menu *MenuHandler,
	division *DivisionHandler,
	roleMenu *RoleMenuHandler,
	rolePermission *RolePermissionHandler,

	// dependencies
	userRepo repository.UserRepository,
	authUC usecase.AuthUsecase,
	jwtSecret string,
) {

	// =====================================================
	// AUTH (PUBLIC)
	// =====================================================
	authGroup := rg.Group("/auth")
	{
		// @Summary Login
		// @Tags Auth
		// @Accept json
		// @Produce json
		// @Router /auth/login [post]
		authGroup.POST(
			"/login",
			middleware.RateLimiter(),
			middleware.ActivityLog("login"),
			auth.Login,
		)

		// @Summary Register
		// @Tags Auth
		// @Accept json
		// @Produce json
		// @Router /auth/register [post]
		authGroup.POST(
			"/register",
			middleware.ActivityLog("register"),
			auth.Register,
		)

		// @Summary Refresh Token
		// @Tags Auth
		// @Accept json
		// @Produce json
		// @Router /auth/refresh [post]
		authGroup.POST(
			"/refresh",
			middleware.ActivityLog("refresh_token"),
			auth.RefreshToken,
		)

		// @Summary Logout
		// @Tags Auth
		// @Security BearerAuth
		// @Router /auth/logout [post]
		authGroup.POST(
			"/logout",
			middleware.JWTAuth(jwtSecret),
			middleware.ActivityLog("logout"),
			auth.Logout,
		)

		// @Summary Forgot Password
		// @Tags Auth
		// @Accept json
		// @Produce json
		// @Router /auth/forgot-password [post]
		authGroup.POST(
			"/forgot-password",
			middleware.ActivityLog("forgot_password"),
			auth.ForgotPassword,
		)
	}

	// =====================================================
	// PROTECTED API
	// =====================================================
	api := rg.Group("/v1")

	api.Use(
		middleware.RateLimiter(),

		// 🔥 AUTO REFRESH TOKEN
		middleware.AutoRefreshToken(authUC),

		// JWT VALIDATION
		middleware.JWTAuth(jwtSecret),

		// USER ACTIVE CHECK
		middleware.IsActiveUser(userRepo),
	)

	// ================= USERS =================
	// @Tags Users
	api.GET(
		"/users",
		middleware.RequirePermission(userRepo, "user.read"),
		middleware.ActivityLog("view_users"),
		user.GetAll,
	)

	api.POST(
		"/users",
		middleware.RequirePermission(userRepo, "user.create"),
		middleware.ActivityLog("create_user"),
		middleware.AuditLog("create", "users"),
		user.Create,
	)

	api.GET(
		"/users/:id",
		middleware.RequirePermission(userRepo, "user.read"),
		user.GetByID,
	)

	api.PUT(
		"/users/:id",
		middleware.RequirePermission(userRepo, "user.update"),
		middleware.ActivityLog("update_user"),
		middleware.AuditLog("update", "users"),
		user.Update,
	)

	api.DELETE(
		"/users/:id",
		middleware.RequirePermission(userRepo, "user.delete"),
		middleware.ActivityLog("delete_user"),
		middleware.AuditLog("delete", "users"),
		user.Delete,
	)

	// ================= ROLES =================
	api.GET(
		"/roles",
		middleware.RequirePermission(userRepo, "role.read"),
		role.GetAll,
	)

	api.POST(
		"/roles",
		middleware.RequirePermission(userRepo, "role.create"),
		middleware.ActivityLog("create_role"),
		middleware.AuditLog("create", "roles"),
		role.Create,
	)

	api.PUT(
		"/roles/:id",
		middleware.RequirePermission(userRepo, "role.update"),
		middleware.ActivityLog("update_role"),
		middleware.AuditLog("update", "roles"),
		role.Update,
	)

	api.DELETE(
		"/roles/:id",
		middleware.RequirePermission(userRepo, "role.delete"),
		middleware.ActivityLog("delete_role"),
		middleware.AuditLog("delete", "roles"),
		role.Delete,
	)

	// ================= DIVISIONS =================
	api.GET(
		"/divisions",
		middleware.RequirePermission(userRepo, "division.read"),
		division.GetAll,
	)

	api.POST(
		"/divisions",
		middleware.RequirePermission(userRepo, "division.create"),
		middleware.ActivityLog("create_division"),
		middleware.AuditLog("create", "divisions"),
		division.Create,
	)

	api.PUT(
		"/divisions/:id",
		middleware.RequirePermission(userRepo, "division.update"),
		middleware.ActivityLog("update_division"),
		middleware.AuditLog("update", "divisions"),
		division.Update,
	)

	api.DELETE(
		"/divisions/:id",
		middleware.RequirePermission(userRepo, "division.delete"),
		middleware.ActivityLog("delete_division"),
		middleware.AuditLog("delete", "divisions"),
		division.Delete,
	)

	// ================= MENUS =================
	api.GET(
		"/menus",
		middleware.RequirePermission(userRepo, "menu.read"),
		menu.GetAll,
	)

	api.POST(
		"/menus",
		middleware.RequirePermission(userRepo, "menu.create"),
		middleware.ActivityLog("create_menu"),
		middleware.AuditLog("create", "menus"),
		menu.Create,
	)

	api.PUT(
		"/menus/:id",
		middleware.RequirePermission(userRepo, "menu.update"),
		middleware.ActivityLog("update_menu"),
		middleware.AuditLog("update", "menus"),
		menu.Update,
	)

	api.DELETE(
		"/menus/:id",
		middleware.RequirePermission(userRepo, "menu.delete"),
		middleware.ActivityLog("delete_menu"),
		middleware.AuditLog("delete", "menus"),
		menu.Delete,
	)

	// ================= ROLE → MENU =================
	api.POST(
		"/roles/menus",
		middleware.RequirePermission(userRepo, "role.assign_menu"),
		middleware.ActivityLog("assign_role_menu"),
		middleware.AuditLog("assign", "role_menu"),
		roleMenu.AssignMenu,
	)

	// ================= ROLE → PERMISSION =================
	api.POST(
		"/roles/permissions",
		middleware.RequirePermission(userRepo, "role.assign_permission"),
		middleware.ActivityLog("assign_role_permission"),
		middleware.AuditLog("assign", "role_permission"),
		rolePermission.AssignPermission,
	)
}
