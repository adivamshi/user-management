package server

import (
	"fmt"
	"user-management-api/routes"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Run(port string) {
	r := gin.Default()
	_ = r.SetTrustedProxies([]string{"127.0.0.1", "localhost", "http://localhost:8081", "0.0.0.0"})
	//corsConfig := cors.DefaultConfig()
	//corsConfig.AllowAllOrigins = true
	r.Use(CORSMiddleware())
	r.Use(gin.Recovery())

	routes.SetupRoutes(r)
	addr := fmt.Sprintf(":%s", port)
	err := r.Run(addr)
	if err != nil {
		return
	}
}
