package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lbrulet/APINIT-GO/middleware"
	"github.com/lbrulet/APINIT-GO/models"
	"github.com/lbrulet/APINIT-GO/mongodb"
	"github.com/lbrulet/APINIT-GO/utils"
)

// RegisterUsersService group every method about the user controller
func RegisterUsersService(route *gin.RouterGroup) {
	route.GET("/", middleware.IsAdmin, func(c *gin.Context) {
		db := mongodb.Database
		if users, err := db.FindAll(); err != nil {
			utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: err.Error()})
		} else {
			utils.SendResponse(c, 200, &models.ResponsePayload{Success: true, Message: users})
		}
	})

	route.GET("/:id", middleware.IsAuthorized, func(c *gin.Context) {
		id := c.Param("id")
		db := mongodb.Database
		if id == "me" {
			id = c.MustGet("id").(string)
		} else {
			if user, err := db.FindByID(c.MustGet("id").(string)); err != nil || user.Admin != true {
				utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Not authorized."})
				return
			}
		}
		if user, err := db.FindByID(id); err != nil {
			utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		} else {
			utils.SendResponse(c, 200, &models.ResponsePayload{Success: true, Message: user})
		}
	})

	route.PUT("/:id", middleware.IsAuthorized, func(c *gin.Context) {
		id := c.Param("id")
		db := mongodb.Database
		if id == "me" {
			id = c.MustGet("id").(string)
		} else {
			if user, err := db.FindByID(c.MustGet("id").(string)); err != nil || user.Admin != true {
				utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Not authorized."})
				return
			}
		}
		if user, err := db.FindByID(id); err != nil {
			utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		} else {
			utils.SendResponse(c, 200, &models.ResponsePayload{Success: true, Message: user})
		}
	})
}
