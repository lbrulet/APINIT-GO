package users

import (
	"github.com/gin-gonic/gin"

	"github.com/lbrulet/APINIT-GO/controllers"
	"github.com/lbrulet/APINIT-GO/middleware"
)

// RegisterUsersService group every method about the user controller
func RegisterUserService(route *gin.RouterGroup) {

	route.GET("/me", middleware.IsAuthorized, controllers.GetMe)
	route.PUT("/me", middleware.IsAuthorized, controllers.PutMe)
	route.DELETE("/me", middleware.IsAuthorized, controllers.DeleteMe)
}
