package routes

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myrachanto/astore/controllers"
)

func StoreApi() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in routes")
	}
	PORT := os.Getenv("PORT")
	key := os.Getenv("EncryptionKey")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) 
	e.Use(middleware.CORS())

	JWTgroup := e.Group("/api/")
	JWTgroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey: []byte(key),
	}))
	//JwtG := e.Group("/users")
	// JwtG.Use(middleware.JWT([]byte(key)))
	// Routes
	e.POST("/register", controllers.UserController.Create)
	e.POST("/login", controllers.UserController.Login)
	JWTgroup.GET("users/logout/:token", controllers.UserController.Logout)
	JWTgroup.GET("users", controllers.UserController.GetAll)
	JWTgroup.GET("users/:id", controllers.UserController.GetOne)
	JWTgroup.PUT("users/:id", controllers.UserController.Update)
	JWTgroup.DELETE("users/:id", controllers.UserController.Delete)
	//e.DELETE("loggoutall/:id", controllers.UserController.DeleteALL) logout all accounts
	///////////category/////////////////////////////	
	JWTgroup.POST("categorys", controllers.CategoryController.Create)
	JWTgroup.GET("categorys", controllers.CategoryController.GetAll)
	JWTgroup.GET("categorys/:id", controllers.CategoryController.GetOne)
	JWTgroup.PUT("categorys/:id", controllers.CategoryController.Update)
	JWTgroup.DELETE("categorys/:id", controllers.CategoryController.Delete)
	///////////majorcategory/////////////////////////////	
	JWTgroup.POST("majorcategory", controllers.MCategoryController.Create)
	JWTgroup.GET("majorcategory", controllers.MCategoryController.GetAll)
	JWTgroup.GET("majorcategory/:id", controllers.MCategoryController.GetOne)
	JWTgroup.PUT("majorcategory/:id", controllers.MCategoryController.Update)
	JWTgroup.DELETE("majorcategory/:id", controllers.MCategoryController.Delete)
	///////////subcategory/////////////////////////////	
	JWTgroup.POST("subcategory", controllers.SubcategoryController.Create)
	JWTgroup.GET("subcategory", controllers.SubcategoryController.GetAll)
	JWTgroup.GET("subcategory/:id", controllers.SubcategoryController.GetOne)
	JWTgroup.PUT("subcategory/:id", controllers.SubcategoryController.Update)
	JWTgroup.DELETE("subcategory/:id", controllers.SubcategoryController.Delete)
	///////////subcategory/////////////////////////////	
	JWTgroup.POST("products", controllers.ProductController.Create)
	JWTgroup.GET("products", controllers.ProductController.GetAll)
	JWTgroup.GET("products/:id", controllers.ProductController.GetOne)
	JWTgroup.PUT("products/:id", controllers.ProductController.Update)
	JWTgroup.DELETE("products/:id", controllers.ProductController.Delete)
	///////////cart/////////////////////////////	
	JWTgroup.POST("cart", controllers.CartController.Create)
	JWTgroup.GET("cart", controllers.CartController.GetAll)
	JWTgroup.GET("cart/:id", controllers.CartController.GetOne)
	JWTgroup.PUT("cart/:id", controllers.CartController.Update)
	JWTgroup.DELETE("cart/:id", controllers.CartController.Delete)
	///////////Invoice/////////////////////////////	
	JWTgroup.POST("invoice", controllers.InvoiceController.Create)
	JWTgroup.GET("invoice", controllers.InvoiceController.GetAll)
	JWTgroup.GET("invoice/:id", controllers.InvoiceController.GetOne)
	JWTgroup.PUT("invoice/:id", controllers.InvoiceController.Update)
	JWTgroup.DELETE("invoice/:id", controllers.InvoiceController.Delete)
	///////////trasanctions/////////////////////////////	
	JWTgroup.POST("trasanctions", controllers.TransactionController.Create)
	JWTgroup.GET("trasanctions", controllers.TransactionController.GetAll)
	JWTgroup.GET("trasanctions/:id", controllers.TransactionController.GetOne)
	JWTgroup.PUT("trasanctions/:id", controllers.TransactionController.Update)
	JWTgroup.DELETE("trasanctions/:id", controllers.TransactionController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
