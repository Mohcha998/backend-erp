package v1

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	rg *gin.RouterGroup,
	auth *AuthHandler,
	user *UserHandler,
	role *RoleHandler,
	menu *MenuHandler,
	division *DivisionHandler,
	roleMenu *RoleMenuHandler,
	rolePermission *RolePermissionHandler,
) {

	// ================= AUTH =================
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/logout", auth.Logout)
		authGroup.POST("/forgot-password", auth.ForgotPassword)
	}

	// ================= API V1 =================
	v1 := rg.Group("/v1")
	{
		// -------- USERS --------
		v1.POST("/users", user.Create)
		v1.GET("/users", user.GetAll)
		v1.GET("/users/:id", user.GetByID)
		v1.PUT("/users/:id", user.Update)
		v1.DELETE("/users/:id", user.Delete)

		// -------- ROLES --------
		v1.POST("/roles", role.Create)
		v1.GET("/roles", role.GetAll)
		v1.GET("/roles/:id", role.GetByID)
		v1.PUT("/roles/:id", role.Update)
		v1.DELETE("/roles/:id", role.Delete)

		// -------- DIVISIONS --------
		v1.POST("/divisions", division.Create)
		v1.GET("/divisions", division.GetAll)
		v1.GET("/divisions/:id", division.GetByID)
		v1.PUT("/divisions/:id", division.Update)
		v1.DELETE("/divisions/:id", division.Delete)

		// -------- MENUS --------
		v1.POST("/menus", menu.Create)
		v1.GET("/menus", menu.GetAll)
		v1.GET("/menus/:id", menu.GetByID)
		v1.PUT("/menus/:id", menu.Update)
		v1.DELETE("/menus/:id", menu.Delete)

		// -------- ROLE → MENU --------
		v1.POST("/roles/menus", roleMenu.AssignMenu)

		// -------- ROLE → PERMISSION --------
		v1.POST("/roles/permissions", rolePermission.AssignPermission)
	}
}
