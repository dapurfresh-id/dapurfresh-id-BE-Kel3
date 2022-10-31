package main

import (
	"fmt"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/controllers"
	"github.com/aldisaputra17/dapur-fresh-id/database"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	contextTimeOut     time.Duration                   = 10 * time.Second
	db                 *gorm.DB                        = database.ConnectDB()
	userRepository     repositories.UserRepository     = repositories.NewUserRepository(db)
	categoryRepository repositories.CategoryRepository = repositories.NewCategoryRepository(db)
	authService        services.AuthService            = services.NewAuthService(userRepository, contextTimeOut)
	jwtService         services.JWTService             = services.NewJWTService()
	categoryService    services.CategoryService        = services.NewCategoryService(categoryRepository, contextTimeOut)
	authController     controllers.AuthController      = controllers.NewAuthController(authService, jwtService)
	categoryController controllers.CategoryController  = controllers.NewCategoryController(categoryService)
)

func main() {
	fmt.Println("Start Server")
	defer database.CloseDatabaseConnection(db)

	r := gin.Default()

	r.Use(cors.Default())

	v1 := r.Group("api")

	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	categoryRoutes := v1.Group("/category")
	{
		categoryRoutes.GET("", categoryController.GetAllCategory)
		categoryRoutes.GET("/:id", categoryController.GetCategoryById)
	}
	r.Run()
}
