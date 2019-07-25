package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lbrulet/APINIT-GO/configs"
	"github.com/lbrulet/APINIT-GO/models"
	"github.com/lbrulet/APINIT-GO/mongodb"
	"github.com/lbrulet/APINIT-GO/utils"
	"github.com/sethvargo/go-password/password"
	"gopkg.in/mgo.v2/bson"
)

// LoginController is a function that handle the login route
func LoginController(c *gin.Context) {
	payload := models.LoginPayload{}
	db := mongodb.Database
	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err == nil {
		if user, err := db.FindByKey("username", payload.Username); err == nil {
			if user.Verified == true {
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
				utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "Account is not verified."})
			}
		} else {
			utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "Account does not exist."})
		}
	} else {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
	}
}

// RegisterController is a function that handle the register route
func RegisterController(c *gin.Context) {
	payload := models.RegisterPayload{}
	db := mongodb.Database
	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err == nil {
		if _, err := db.FindByKey("username", payload.Username); err != nil {
			if _, err := db.FindByKey("email", payload.Email); err != nil {
				var person models.User
				person.ID = bson.NewObjectId()
				person.Username = payload.Username
				person.Password = payload.Password
				person.Email = payload.Email
				person.AuthMethod = models.LOCAL
				if err := db.Insert(person); err != nil {
					utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Database unavailable."})
				} else {
					if token, err := utils.CreateToken(person, time.Now().Add(time.Hour*configs.Config.AccessTokenValidityTime).Unix()); err != nil {
						utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: err.Error()})
					} else {
						if pwd, err := os.Getwd(); err != nil {
							panic(err)
						} else {
							utils.SendMail(person, models.Template{
								Email:        person.Email,
								Username:     person.Username,
								ConfirmEmail: configs.Config.MailConfirmationLink + "?token=" + token,
							}, pwd+"/templates/welcome.html")
						}
					}
					utils.SendResponse(c, http.StatusCreated, &models.ResponsePayload{Success: true, Message: "Account created."})
				}
			} else {
				utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Email already exist."})
			}
		} else {
			utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Username already exist."})
		}
	} else {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
	}
}

// RecoveryController is a function that handle the recovery password route
func RecoveryController(c *gin.Context) {
	payload := models.RecoveryPayload{}
	db := mongodb.Database
	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err == nil {
		if user, err := db.FindByKey("email", payload.Email); err == nil {
			if res, err := password.Generate(7, 2, 2, false, false); err != nil {
				c.Redirect(http.StatusMovedPermanently, configs.Config.MailSuccessRedirect)
			} else {
				if pwd, err := os.Getwd(); err != nil {
					panic(err)
				} else {
					utils.SendMail(user, models.TemplateRecovery{
						Email:    user.Email,
						Username: user.Username,
						Password: res,
					}, pwd+"/templates/recovery.html")
				}
				c.Redirect(http.StatusMovedPermanently, configs.Config.MailSuccessRedirect)
			}
		} else {
			c.Redirect(http.StatusMovedPermanently, configs.Config.MailFailedRedirect)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, configs.Config.MailFailedRedirect)
	}
}

// ConfirmAccountController is a function that hundle the confirm account route
func ConfirmAccountController(c *gin.Context) {
	token := c.Query("token")
	db := mongodb.Database
	if claims, err := utils.ExtractClaims(token); err != nil {
		c.Redirect(http.StatusMovedPermanently, configs.Config.MailFailedRedirect)
	} else {
		if user, err := db.FindByID(claims.ID); err != nil {
			c.Redirect(http.StatusMovedPermanently, configs.Config.MailFailedRedirect)
		} else {
			user.Verified = true
			if err := db.Update(user); err != nil {
				c.Redirect(http.StatusMovedPermanently, configs.Config.MailFailedRedirect)
			} else {
				c.Redirect(http.StatusMovedPermanently, configs.Config.MailSuccessRedirect)
			}
		}
	}
}

// SecretController is a function that hundle a secret way to the api
func SecretController(c *gin.Context) {
	c.JSON(200, &models.ResponsePayload{Success: true, Message: c.MustGet("id").(string)})
}
