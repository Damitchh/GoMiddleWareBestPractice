package router

import (
	"Hacktiv10JWT/controllers"
	"Hacktiv10JWT/middlewares"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", controllers.GetProducts)

		// Restrict PUT and DELETE routes to only admin users
		adminRouter := productRouter.Group("")
		adminRouter.Use(middlewares.Authentication())
		adminRouter.Use(middlewares.AdminOnly())
		adminRouter.GET("/all", controllers.GetAllProducts)
		adminRouter.GET("/:ID", controllers.GetProductbyID)
		adminRouter.PUT("/:ID", controllers.UpdateProduct)
		adminRouter.DELETE("/:ID", controllers.DeleteProduct)

	}

	return r
}
