// @title Auth Service API
// @version 1.0
// @description Authentication & Authorization Service
// @termsOfService https://yourcompany.com/terms

// @contact.name API Support
// @contact.email support@yourcompany.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8081
// @BasePath /api
// @schemes http

// 🔐 JWT AUTH
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main
import "fmt"

import (
	_ "auth-service/docs" // ← WAJIB untuk swagger
	"auth-service/internal/server"
)

func main() {
	fmt.Println("Starting the server...") 
	server.Run()
	fmt.Println("Server has started.")
}
