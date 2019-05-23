package routes

import (
	"fmt"

	"github.com/Zulbukharov/webhook_service/handlers"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes ...
func InitializeRoutes(router *gin.Engine) {
	router.POST("/article", handlers.Article)
	fmt.Println("routes initialized")
}
