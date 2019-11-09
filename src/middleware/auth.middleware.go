package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/lbrulet/APINIT-GO/src/database"
	"github.com/lbrulet/APINIT-GO/src/models"
	"github.com/lbrulet/APINIT-GO/src/utils"
)

// IsAuthorized is a middleware that check if the token is valid
// id := c.MustGet("id").(string) to get the setting value by gin-gonic
func IsAuthorized(c *gin.Context) {
	if c.Request.Header["Authorization"] == nil {
		utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Authorization is missing"})
		c.Abort()
		return
	}

	token := c.Request.Header["Authorization"][0][7:len(c.Request.Header["Authorization"][0])]
	claims := &models.Claims{}

	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("pingouin123"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Not Authorized"})
		} else {
			utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: err.Error()})
		}
		c.Abort()
		return
	} else if !tk.Valid {
		utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Not Authorized"})
		c.Abort()
		return
	}
	if _, err := database.Database.GetUserByID(claims.ID); err != nil {
		utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Not Authorized"})
		c.Abort()
		return
	}
	c.Set("id", claims.ID)
	c.Next()
}

// IsAdmin is a middleware that check if the token is valid and if the client is an administrator
// id := c.MustGet("id").(string) to get the setting value by gin-gonic
func IsAdmin(c *gin.Context) {
	if c.Request.Header["Authorization"] == nil {
		utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Admin reserved action"})
		c.Abort()
		return
	}

	token := c.Request.Header["Authorization"][0][7:len(c.Request.Header["Authorization"][0])]
	claims := &models.Claims{}

	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("pingouin123"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Admin reserved action"})
		} else {
			utils.SendResponse(c, http.StatusBadRequest, &models.ResponsePayload{Success: false, Message: err.Error()})
		}
		c.Abort()
		return
	} else if !tk.Valid {
		utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Admin reserved action"})
		c.Abort()
		return
	}

	if user, err := database.Database.GetUserByID(claims.ID); err == nil && user.Admin {
		c.Set("id", claims.ID)
		c.Next()
	} else {
		utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Admin reserved action"})
		c.Abort()
		return
	}
}
