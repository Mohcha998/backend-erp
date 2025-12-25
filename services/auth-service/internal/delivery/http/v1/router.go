package v1

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	rg *gin.RouterGroup,
	auth *AuthHandler,
	user *UserHandler,
	role *RoleHandler,
	menu *MenuHandler,
	division *DivisionHandler,
	permission *PermissionHandler,
) {
	// AUTH
	RegisterAuthRoutes(rg, auth)

	v1 := rg.Group("/v1")
	{
		v1.POST("/users", user.Create)
		v1.GET("/users", user.GetAll)

		v1.POST("/roles", role.Create)
		v1.GET("/roles", role.GetAll)

		v1.POST("/divisions", division.Create)
		v1.GET("/divisions", division.GetAll)
		v1.DELETE("/divisions/:id", division.Delete)

		v1.POST("/menus", menu.Create)
		v1.GET("/menus", menu.GetAll)

		v1.POST("/permissions", permission.Create)
		v1.GET("/permissions", permission.GetAll)
	}
}
