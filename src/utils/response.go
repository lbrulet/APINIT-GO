package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/APINIT-GO/src/models"
)

// SendResponse is used to respond to the client
func SendResponse(c *gin.Context, code int, response *models.ResponsePayload) {
	c.JSON(code, gin.H{"success": response.Success, "message": response.Message})
}

// SendLoginResponse is used to respond to the client
func SendLoginResponse(c *gin.Context, code int, response *models.LoginResponsePayload) {
	c.JSON(code, gin.H{"success": response.Success, "message": response.Message, "token": response.Token, "refresh-token": response.RefreshToken, "user": response.User})
}

// SendRegisterResponse is used to respond to the client
func SendRegisterResponse(c *gin.Context, code int, response *models.RegisterResponsePayload) {
	c.JSON(code, gin.H{"success": response.Success, "message": response.Message, "confirm_account_token": response.Token, "user": response.User})
}
