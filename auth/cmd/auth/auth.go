package main

import (
	"saas/auth/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.AuthRoutes(r)

	r.Run(":8081") // Run on localhost:8080
}
