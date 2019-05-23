package main

import (
	"github.com/Zulbukharov/webhook_service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitializeRoutes(router)
	router.Run(":8080")
}
