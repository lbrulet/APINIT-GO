package models

import "time"

// Configuration struct is used to load configuration from file or environnement
type Configuration struct {
	Port                     string        `json:"app_port" env:"APP_PORT"`
	DatabaseName             string        `json:"database_name" env:"DATABASE_NAME"`
	DatabaseCollection       string        `json:"database_collection" env:"DATABASE_COLLECTION"`
	DatabaseHost             string        `json:"database_host" env:"DATABASE_HOST"`
	AccessTokenValidityTime  time.Duration `json:"access_token_validity_time" env:"ACCESS_TOKEN_VALIDITY_TIME"`
	RefreshTokenValidityTime time.Duration `json:"refresh_token_validity_time" env:"REFRESH_TOKEN_VALIDITY_TIME"`
}
