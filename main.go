package main

import (
	"fmt"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/controllers"
	"github.com/aldisaputra17/dapur-fresh-id/database"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	contextTimeOut time.Duration               = 10 * time.Second
	db             *gorm.DB                    = database.ConnectDB()
	userRepository repositories.UserRepository = repositories.NewUserRepository(db)
	authService    services.AuthService        = services.NewAuthService(userRepository, contextTimeOut)
	jwtService     services.JWTService         = services.NewJWTService()
	authController controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
)

func main() {
	fmt.Println("Start Server")
	defer database.CloseDatabaseConnection(db)

	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
