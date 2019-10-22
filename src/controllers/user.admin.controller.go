package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lbrulet/APINIT-GO/src/database"
	"github.com/lbrulet/APINIT-GO/src/models"
	"github.com/lbrulet/APINIT-GO/src/utils"
	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers get all users
func GetAllUsers(c *gin.Context) {
	users, err := database.Database.GetAllUsers()
	if err != nil {
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Database unavailable."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: users})
}

// GetUserByID get a user by id
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: err.Error()})
		return
	}
	user, err := database.Database.GetUserByID(id)
	if err != nil {
		fmt.Println(err, id)
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: user})
}

// UpdateUserByID update a user by id
func UpdateUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	payload := models.UserUpdate{}

	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: err.Error()})
		return
	}

	// Check existing user
	user, err := database.Database.GetUserByID(id)
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
		if count, err := database.Database.CountUserByKey("email", payload.Email); err != nil {
			utils.SendResponse(c, http.StatusInternalServerError, &models.ResponsePayload{Success: false, Message: err.Error()})
			return
		} else if count > 0 {
			utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Email already exist."})
			return
		}
		user.Email = payload.Email
	}

	// Check existing username
	if len(payload.Username) > 0 && user.Username != payload.Username {
		if count, err := database.Database.CountUserByKey("username", payload.Username); err != nil {
			utils.SendResponse(c, http.StatusInternalServerError, &models.ResponsePayload{Success: false, Message: err.Error()})
			return
		} else if count > 0 {
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
		user.Password = string(hash)
	}

	user.Admin = payload.Admin
	user.Verified = payload.Verified

	if err := user.Update(database.Database.DB); err != nil {
		fmt.Println(err)
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Database unavailable."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: "Success update!"})
}

// DeleteUser delete user by id
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: err.Error()})
		return
	}

	// Check existing user
	user, err := database.Database.GetUserByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		return
	}

	// Delete user
	if err := user.Delete(database.Database.DB); err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "Database unvailable."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: "Deleted with success."})
}
