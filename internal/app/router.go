package app

import (
	"core-api/internal/interfaces/handlers"
	"core-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handlers.AuthHandler, accountHandler *handlers.AccountHandler) *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/login", authHandler.Login)
		apiV1.POST("/register", authHandler.Register)
		router.POST("/users/accounts", accountHandler.AssociateUserToAccount)
		router.POST("/accounts", accountHandler.CreateAccount)

		authenticatedRoutes := apiV1.Group("/")
		authenticatedRoutes.Use(middleware.JWTAuthMiddleware())
		{
			authenticatedRoutes.GET("/securedata", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Accès sécurisé aux données réussi",
				})
			})
		}
	}

	return router
}
