package main

import (
	"fmt"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/controllers"
	"github.com/aldisaputra17/dapur-fresh-id/database"
	"github.com/aldisaputra17/dapur-fresh-id/middleware"
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
	cartRepository     repositories.CartRepository     = repositories.NewCartRepository(db)
	authService        services.AuthService            = services.NewAuthService(userRepository, contextTimeOut)
	jwtService         services.JWTService             = services.NewJWTService()
	categoryService    services.CategoryService        = services.NewCategoryService(categoryRepository, contextTimeOut)
	cartService        services.CartService            = services.NewCartService(cartRepository, contextTimeOut)
	userService        services.UserService            = services.NewUserService(userRepository, contextTimeOut)
	authController     controllers.AuthController      = controllers.NewAuthController(authService, jwtService)
	categoryController controllers.CategoryController  = controllers.NewCategoryController(categoryService)
	cartController     controllers.CartController      = controllers.NewCartController(cartService, jwtService)
	userController     controllers.UserController      = controllers.NewUserController(userService, jwtService)
)

func main() {
	fmt.Println("Start Server")
	defer database.CloseDatabaseConnection(db)

	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})

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
	cartRoutes := v1.Group("/cart", middleware.AuthorizeJWT(jwtService))
	{
		cartRoutes.POST("", cartController.AddCart)
	}
	userRoutes := v1.Group("/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.PUT("", userController.Update)
	}
	r.Run()
}
