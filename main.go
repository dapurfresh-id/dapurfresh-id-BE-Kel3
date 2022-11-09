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
	contextTimeOut     time.Duration                   = 10 * time.Second
	db                 *gorm.DB                        = database.ConnectDB()
	userRepository     repositories.UserRepository     = repositories.NewUserRepository(db)
	categoryRepository repositories.CategoryRepository = repositories.NewCategoryRepository(db)
	productRepository  repositories.ProductRepository  = repositories.NewProductRepository(db)
	authService        services.AuthService            = services.NewAuthService(userRepository, contextTimeOut)
	jwtService         services.JWTService             = services.NewJWTService()
	categoryService    services.CategoryService        = services.NewCategoryService(categoryRepository, contextTimeOut)
	productService     services.ProductService         = services.NewProductService(productRepository, contextTimeOut)
	authController     controllers.AuthController      = controllers.NewAuthController(authService, jwtService)
	categoryController controllers.CategoryController  = controllers.NewCategoryController(categoryService)
	productController  controllers.ProductController   = controllers.NewProductController(productService)
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
	productRoutes := r.Group("/product")
	{
		productRoutes.GET("/", productController.GetAllProduct)
		productRoutes.GET("/:id", productController.GetProductById)
		productRoutes.GET("/search", productController.PaginationProduct)
		productRoutes.GET("/category", productController.GetProductByCategory)
		productRoutes.GET("/limit", productController.GetLimitProduct)
		nameProductRoutes := productRoutes.Group("/name")
		{
			nameProductRoutes.GET("equal", productController.GetProductByNameEqual)
			nameProductRoutes.GET("contains", productController.GetProductByNameContains)
			nameProductRoutes.GET("like", productController.GetProductByNameLike)
		}
	}
	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.GET("/", categoryController.GetAllCategory)
		categoryRoutes.GET("/:id", categoryController.GetCategoryById)
	}
	r.Run()
}
