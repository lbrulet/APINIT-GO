package users

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/APINIT-GO/src/controllers"
	"github.com/lbrulet/APINIT-GO/src/middleware"
)

func RegisterAdminUserService(route *gin.RouterGroup) {

	route.GET("/", middleware.IsAdmin, controllers.GetAllUsers)

	route.GET("/:id", middleware.IsAdmin, controllers.GetUserByID)

	route.PUT("/:id", middleware.IsAuthorized, controllers.UpdateUserByID)

	route.DELETE("/:id", middleware.IsAuthorized, controllers.DeleteUser)
}
