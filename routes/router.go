package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/APINIT-GO/routes/authentication"
	"github.com/lbrulet/APINIT-GO/routes/users"
)

// CORS allow request from outside
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, authorization, content-type")
		c.Header("Content-Type", "application/json")
	}
}

// InitRouter return a http server
func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(CORS(), gin.Logger(), gin.Recovery())

	api := router.Group("/api")

	auth := api.Group("/auth")

	authentication.RegisterAuthService(auth)

	user := api.Group("/users")

	users.RegisterUsersService(user)

	return router
}
