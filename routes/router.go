package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTION", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	auth := api.Group("/auth")

	authentication.RegisterAuthService(auth)

	user := api.Group("/users")

	users.RegisterUsersService(user)

	return router
}
