package authentication

import (
	"github.com/gin-gonic/gin"

	"github.com/lbrulet/APINIT-GO/src/controllers"
	"github.com/lbrulet/APINIT-GO/src/middleware"
)

// RegisterAuthService add route handler from the authentication
func RegisterAuthService(route *gin.RouterGroup) {
	route.GET("/secret", middleware.IsAuthorized, controllers.SecretController)

	route.GET("/confirm-account", controllers.ConfirmAccountController)

	route.POST("/recovery", controllers.RecoveryController)

	route.POST("/login", controllers.LoginController)

	route.POST("/register", controllers.RegisterController)
}
