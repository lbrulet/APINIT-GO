package routes

import (
	"github.com/gin-gonic/gin"

	_ "github.com/lbrulet/APINIT-GO/docs"
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

	apiRouter := router.Group("/api")

	authRouter := apiRouter.Group("/auth")
	authentication.RegisterAuthService(authRouter)

	userRouter := apiRouter.Group("/user")
	users.RegisterUserService(userRouter)

	adminUserRouter := apiRouter.Group("/admin/user")
	users.RegisterAdminUserService(adminUserRouter)

	return router
}
