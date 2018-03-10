package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware Middleware to configure CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}

// routerEngine Configure routes
func routerEngine() *gin.Engine {
	r := gin.New()

	// Global middlewares
	r.Use(CORSMiddleware())

	// Necessary to CORS
	r.OPTIONS("api/v1/*cors", func(c *gin.Context) {
		// Empty 200 response
	})

	return r
}
