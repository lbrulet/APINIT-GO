package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lbrulet/APINIT-GO/models"
	"github.com/lbrulet/APINIT-GO/mongodb"
	"github.com/lbrulet/APINIT-GO/utils"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context) {
	db := mongodb.Database

	users, err := db.FindAll()
	if err != nil {
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Database unavailable."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: users})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	db := mongodb.Database

	user, err := db.FindByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: user})
}

func UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	db := mongodb.Database
	payload := models.UserUpdate{}

	// Check existing user
	user, err := db.FindByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		return
	}

	// Load payload
	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
		return
	}

	// Check existing email
	if len(payload.Email) > 0 && user.Email != payload.Email {
		if _, err := db.FindByKey("email", payload.Email); err == nil {
			utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Email already exist."})
			return
		}
		user.Email = payload.Email
	}

	// Check existing username
	if len(payload.Username) > 0 && user.Username != payload.Username {
		if _, err := db.FindByKey("username", payload.Username); err == nil {
			utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Username already exist."})
			return
		}
		user.Username = payload.Username
	}

	// Bcrypt new password
	if len(payload.Password) > 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
		if err != nil {
			utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Hash password unavailable."})
			return
		}
		user.Password = hash
	}

	user.Admin = payload.Admin
	user.Verified = payload.Verified

	if err := db.Update(user); err != nil {
		fmt.Println(err)
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Database unavailable."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: "Success update!"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	db := mongodb.Database

	// Check existing user
	user, err := db.FindByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		return
	}

	// Delete user
	if err := db.Delete(user); err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "Database unvailable."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: "Deleted with success."})
}
