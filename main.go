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
	productRepository  repositories.ProductRepository  = repositories.NewProductRepository(db)
	orderRepository    repositories.OrderRepository    = repositories.NewOrderRepository(db)
	authService        services.AuthService            = services.NewAuthService(userRepository, contextTimeOut)
	jwtService         services.JWTService             = services.NewJWTService()
	categoryService    services.CategoryService        = services.NewCategoryService(categoryRepository, contextTimeOut)
	cartService        services.CartService            = services.NewCartService(cartRepository, contextTimeOut)
	userService        services.UserService            = services.NewUserService(userRepository, contextTimeOut)
	productService     services.ProductService         = services.NewProductService(productRepository, contextTimeOut)
	imgService         services.ImageService           = services.NewImage()
	orderService       services.OrderService           = services.NewOrderService(orderRepository, contextTimeOut)
	authController     controllers.AuthController      = controllers.NewAuthController(authService, jwtService)
	categoryController controllers.CategoryController  = controllers.NewCategoryController(categoryService)
	cartController     controllers.CartController      = controllers.NewCartController(cartService, jwtService)
	userController     controllers.UserController      = controllers.NewUserController(userService, jwtService)
	productController  controllers.ProductController   = controllers.NewProductController(productService)
	imgController      controllers.ImageController     = controllers.NewImgController(imgService, db)
	orderController    controllers.OrderController     = controllers.NewOrderController(orderService, jwtService)
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("Start Server")
	defer database.CloseDatabaseConnection(db)

	r := gin.Default()

	r.Use(cors.Default())

	api := r.Group("api")

	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	categoryRoutes := api.Group("/category")
	{
		categoryRoutes.GET("", categoryController.GetAllCategory)
		categoryRoutes.GET("/:id", categoryController.GetCategoryById)
		categoryRoutes.POST("", categoryController.CreateCategory)
	}
	cartRoutes := api.Group("/cart", middleware.AuthorizeJWT(jwtService))
	{
		cartRoutes.POST("", cartController.AddCart)
		cartRoutes.GET("", cartController.GetCarts)
		cartRoutes.DELETE("/:id", cartController.Delete)
		cartRoutes.PUT("", cartController.Update)
	}
	userRoutes := api.Group("/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.POST("", userController.Update)
	}
	prodRoutes := api.Group("/product")
	{
		prodRoutes.POST("", productController.Create)
	}
	imgRoutes := api.Group("/img")
	{
		imgRoutes.POST("", imgController.Create)
	}
	orderRoutes := api.Group("/checkout", middleware.AuthorizeJWT(jwtService))
	{
		orderRoutes.POST("", orderController.Create)
		orderRoutes.GET("", orderController.GetOrder)
		orderRoutes.GET("/:id", orderController.GetDetail)
		orderRoutes.PATCH("", orderController.PatchStatus)
	}
	// seeder.DBSeed(db)
	r.Run()
}
