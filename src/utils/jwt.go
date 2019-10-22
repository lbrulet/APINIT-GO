package utils

import (
	"errors"

	"github.com/lbrulet/APINIT-GO/src/models"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create a jwt token depending to the user's id
func CreateToken(user *models.User, ExpiresAt int64) (string, error) {

	claims := &models.Claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("pingouin123"))
}

// ExtractClaims from the header and return it
func ExtractClaims(token string) (*models.Claims, error) {
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
