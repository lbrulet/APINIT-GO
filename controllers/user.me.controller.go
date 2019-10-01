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

// GetUser is a function that get the user identity
func GetUser(c *gin.Context) {
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
		utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: user})
	}
}

// UpdateUser is a function that get the user identity
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	db := mongodb.Database
	payload := models.UserUpdate{}

	userAdmin, err := db.FindByID(c.MustGet("id").(string))
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
		return
	}
	if id == "me" {
		id = c.MustGet("id").(string)
	} else {
		if !userAdmin.Admin {
			utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Not authorized."})
			return
		}
	}
	user, err := db.FindByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		return
	}
	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: "Bad request."})
		return
	}
	if len(payload.Email) > 0 {
		if _, err := db.FindByKey("email", payload.Email); err == nil {
			utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Email already exist."})
			return
		}
		user.Email = payload.Email
	}
	if len(payload.Username) > 0 {
		if _, err := db.FindByKey("username", payload.Username); err == nil {
			utils.SendResponse(c, http.StatusConflict, &models.ResponsePayload{Success: false, Message: "Username already exist."})
			return
		}
		user.Username = payload.Username
	}
	if len(payload.Password) > 0 {
		if userAdmin.Admin {
			utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Admin can not modify a user's password"})
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
		if err != nil {
			utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Hash password unavailable."})
			return
		}
		user.Password = hash
	}

	if !userAdmin.Admin {
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Not authorized."})
		return
	}
	user.Admin = payload.Admin

	if !userAdmin.Admin {
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Not authorized."})
		return
	}
	user.Verified = payload.Verified

	if err := db.Update(user); err != nil {
		fmt.Println(err)
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Database unavailable."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: "Success update!"})
}

func GetMe(c *gin.Context) {
	db := mongodb.Database
	id := c.MustGet("id").(string)

	user, err := db.FindByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: user})
}

func PutMe(c *gin.Context) {
	db := mongodb.Database
	id := c.MustGet("id").(string)
	payload := models.UserUpdate{}

	// Check existing user
	user, err := db.FindByID(id)
	if err != nil {
		utils.SendResponse(c, http.StatusNotFound, &models.ResponsePayload{Success: false, Message: "User not found."})
		return
	}

	// load payload
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

	// Bcypt new password
	if len(payload.Password) > 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
		if err != nil {
			utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Hash password unavailable."})
			return
		}
		user.Password = hash
	}

	// Update
	if err := db.Update(user); err != nil {
		utils.SendResponse(c, http.StatusServiceUnavailable, &models.ResponsePayload{Success: false, Message: "Database unavailable."})
		return
	}
	utils.SendResponse(c, http.StatusOK, &models.ResponsePayload{Success: true, Message: "Success update!"})
}

func DeleteMe(c *gin.Context) {
	id := c.MustGet("id").(string)
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
