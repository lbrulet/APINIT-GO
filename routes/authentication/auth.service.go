package authentication

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/lbrulet/APINIT-GO/configs"
	"github.com/lbrulet/APINIT-GO/middleware"
	models "github.com/lbrulet/APINIT-GO/models"
	"github.com/lbrulet/APINIT-GO/mongodb"
	"github.com/lbrulet/APINIT-GO/utils"
)

// RegisterAuthService add route handler from the authentication
func RegisterAuthService(route *gin.RouterGroup) {
	route.Use()

	route.GET("/secret", middleware.IsAuthorized, func(c *gin.Context) {
		c.JSON(200, &models.ResponsePayload{Success: true, Message: c.MustGet("id").(string)})
	})

	route.POST("/login", func(c *gin.Context) {
		payload := models.LoginPayload{}
		db := mongodb.Database
		if err := c.ShouldBindBodyWith(&payload, binding.JSON); err == nil {
			if user, err := db.FindByKey("username", payload.Username); err == nil {
				fmt.Println(user.ID)
				if token, err := utils.CreateToken(user, time.Now().Add(time.Hour*configs.Config.AccessTokenValidityTime).Unix()); err != nil {
					utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: err.Error()})
				} else {
					if refresh, err := utils.CreateToken(user, time.Now().Add(time.Hour*configs.Config.RefreshTokenValidityTime).Unix()); err != nil {
						utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: err.Error()})
					} else {
						utils.SendLoginResponse(c, http.StatusOK, &models.LoginResponsePayload{Success: true, Message: "You are logged in.", Token: token, RefreshToken: refresh})
					}
				}
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
