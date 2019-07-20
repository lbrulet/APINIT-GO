package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/lbrulet/APINIT-GO/models"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create a jwt token depending to the user's id
func CreateToken(user models.User, ExpiresAt int64) (string, error) {

	claims := &models.Claims{
		ID: user.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("pingouin123"))
}

// ExtractClaims from the header and return it
func ExtractClaims(c *gin.Context) (*models.Claims, error) {
	if c.Request.Header["Authorization"] != nil {
		token := c.Request.Header["Authorization"][0][7:len(c.Request.Header["Authorization"][0])]
		claims := &models.Claims{}

		tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("pingouin123"), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return nil, errors.New("Not Authorized")
			}
			return nil, err
		}
		if !tk.Valid {
			return nil, errors.New("Not Authorized")
		}
		return claims, nil
	}
	return nil, errors.New("Authorization is missing")
}
