package middleware

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/APINIT-GO/models"
	"github.com/lbrulet/APINIT-GO/mongodb"
	"github.com/lbrulet/APINIT-GO/utils"
)

type testHeader struct {
	Token string `header:"Authorization"`
}

// IsAuthorized is a middleware that check if the token is valid
// id := c.MustGet("id").(string) to get the setting value by gin-gonic
func IsAuthorized(c *gin.Context) {
	if c.Request.Header["Authorization"] != nil {
		users := mongodb.Database
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
			return
		} else if !tk.Valid {
			utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Not Authorized"})
			return
		}
		if _, err := users.FindByID(claims.ID); err == nil {
			c.Set("id", claims.ID)
			c.Next()
		} else {
			utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Not Authorized"})
		}
		return
	}
	utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Authorization is missing"})
}

// IsAdmin is a middleware that check if the token is valid and if the client is an administrator
// id := c.MustGet("id").(string) to get the setting value by gin-gonic
func IsAdmin(c *gin.Context) {
	if c.Request.Header["Authorization"] != nil {
		users := mongodb.Database
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
			return
		} else if !tk.Valid {
			utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Admin reserved action"})
			return
		}
		if user, err := users.FindByID(claims.ID); err == nil && user.Admin == true {
			c.Set("id", claims.ID)
			c.Next()
		} else {
			utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Admin reserved action"})
		}
		return
	}
	utils.SendResponse(c, http.StatusUnauthorized, &models.ResponsePayload{Success: false, Message: "Admin reserved action"})
}
