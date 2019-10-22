package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claims is the jwt struct
type Claims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
