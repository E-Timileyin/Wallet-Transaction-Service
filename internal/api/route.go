package api

import (
	"wallet-service/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Group routes by resource
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", handlers.GetUsers)    // GET /users
		userRoutes.POST("", handlers.CreateUser) // POST /users
		userRoutes.GET("/:id", handlers.GetUserById)
	}

	// walletRoutes := r.Group("/wallets")
	// {
	//     // Define wallet endpoints later
	// }

	// transactionRoutes := r.Group("/transactions")
	// {
	//     // Define transaction endpoints later
	// }

	return r
}
