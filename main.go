package main

import (
	"fmt"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/controllers"
	"github.com/aldisaputra17/dapur-fresh-id/database"
	"github.com/aldisaputra17/dapur-fresh-id/middleware"
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
	userService    services.UserService        = services.NewUserService(userRepository)
	jwtService     services.JWTService         = services.NewJWTService()
	authController controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController  = controllers.NewUserController(userService, jwtService)
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

	userRoutes := r.Group("/", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("show", userController.GetUser)
		userRoutes.PUT("update:id", userController.UpdateUser)
	}
	r.Run()
}
