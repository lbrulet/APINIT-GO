package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claims is the jwt struct
type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
