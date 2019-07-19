package authentication

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	models "github.com/lbrulet/APINIT-GO/models"
	"github.com/lbrulet/APINIT-GO/mongodb"
	"github.com/lbrulet/APINIT-GO/utils"
)

// RegisterAuthService add route handler from the authentication
func RegisterAuthService(route *gin.RouterGroup) {
	route.Use()

	route.POST("/login", func(c *gin.Context) {
		payload := models.LoginPayload{}
		db := mongodb.Database
		if err := c.ShouldBindBodyWith(&payload, binding.JSON); err == nil {
			user, _ := db.FindByKey("username", payload.Username)
			fmt.Println(user)
			if _, err := db.FindByKey("username", payload.Username); err == nil {
				utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: "You have logged in."})
			} else {
				utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "Account does not exist."})
			}
		} else {
			utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
		}
	})

	route.POST("/register", func(c *gin.Context) {
		payload := models.RegisterPayload{}
		db := mongodb.Database
		if err := c.ShouldBindBodyWith(&payload, binding.JSON); err == nil {
			if _, err := db.FindByKey("username", payload.Username); err != nil {
				if _, err := db.FindByKey("email", payload.Email); err != nil {
					var person models.User
					person.Username = payload.Username
					person.Password = payload.Password
					person.Email = payload.Email
					person.AuthMethod = models.LOCAL
					if err := db.Insert(person); err != nil {
						utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Database unavailable."})
					} else {
						utils.SendResponse(c, http.StatusCreated, &models.ResponsePayload{Success: true, Message: "Account created."})
					}
				} else {
					utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Account already exist."})
				}
			} else {
				utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Account already exist."})
			}
		} else {
			utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
		}
	})
}
