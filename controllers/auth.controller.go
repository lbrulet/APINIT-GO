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
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// LoginController is a function that handle the login route
func LoginController(c *gin.Context) {
	payload := models.LoginPayload{}
	db := mongodb.Database

	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
		return
	}

	user, err := db.FindByKey("username", payload.Username)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "Username or password invalid."})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(payload.Password)); err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "Username or password invalid."})
		return
	}

	if !user.Verified {
		utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Account is not verified."})
		return
	}

	token, err := utils.CreateToken(user, time.Now().Add(time.Hour*configs.Config.AccessTokenValidityTime).Unix())
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
		return
	}

	refresh, err := utils.CreateToken(user, time.Now().Add(time.Hour*configs.Config.RefreshTokenValidityTime).Unix())
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
	}

	utils.SendLoginResponse(c, http.StatusOK, &models.LoginResponsePayload{Success: true, Message: "You are logged in.", Token: token, RefreshToken: refresh, User: user})
}

// RegisterController is a function that handle the register route
func RegisterController(c *gin.Context) {
	payload := models.RegisterPayload{}
	db := mongodb.Database

	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
		return
	}

	if _, err := db.FindByKey("username", payload.Username); err == nil {
		utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Username already exist."})
		return
	}

	if _, err := db.FindByKey("email", payload.Email); err == nil {
		utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Email already exist."})
		return
	}

	var person models.User
	person.ID = bson.NewObjectId()
	person.Username = payload.Username
	person.Email = payload.Email
	person.AuthMethod = models.LOCAL

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Hash password unavailable."})
		return
	}
	person.Password = hash

	if err := db.Insert(person); err != nil {
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Database unavailable."})
		return
	}

	if token, err := utils.CreateToken(person, time.Now().Add(time.Hour*configs.Config.AccessTokenValidityTime).Unix()); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
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
		utils.SendRegisterResponse(c, http.StatusCreated, &models.RegisterResponsePayload{Success: true, Message: "Account created.", Token: token, User: person})
	}
}

// RecoveryController is a function that handle the recovery password route
func RecoveryController(c *gin.Context) {
	payload := models.RecoveryPayload{}
	db := mongodb.Database
	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err == nil {
		if user, err := db.FindByKey("email", payload.Email); err == nil {
			if res, err := password.Generate(7, 2, 2, false, false); err != nil {
				utils.SendResponse(c, http.StatusInternalServerError, &models.ResponsePayload{Success: false, Message: "Bad request."})
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
				utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: "Recovery email sended."})
			}
		} else {
			utils.SendResponse(c, http.StatusInternalServerError, &models.ResponsePayload{Success: false, Message: "Bad request."})
		}
	} else {
		utils.SendResponse(c, http.StatusInternalServerError, &models.ResponsePayload{Success: false, Message: "Bad request."})
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
