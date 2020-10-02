package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kayalova/auth-service/handler"
)

// InitRoutes initiates all routes
func InitRoutes() *gin.Engine {
	router := gin.Default()

	g := router.Group("/api/auth")
	g.GET("/getTokens", handler.GenerateTokens)
	g.PUT("/updateTokens", handler.UpdateTokens)
	g.DELETE("/deleteRefreshToken", handler.DeleteRefreshToken)

	return router
}

// Run server
func Run(r *gin.Engine) {
	r.Run()
}
