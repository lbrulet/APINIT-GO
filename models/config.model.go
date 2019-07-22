package models

import "time"

// Configuration struct is used to load configuration from file or environnement
type Configuration struct {
	Port                 string `json:"app_port" env:"APP_PORT"`
	DatabaseName         string `json:"database_name" env:"DATABASE_NAME"`
	DatabaseCollection   string `json:"database_collection" env:"DATABASE_COLLECTION"`
	DatabaseHost         string `json:"database_host" env:"DATABASE_HOST"`
	MailFrom             string `json:"mail_from" env:"MAIL_FROM"`
	MailSubject          string `json:"mail_subject" env:"MAIL_SUBJECT"`
	MailConfirmationLink string `json:"mail_confirmation" env:"MAIL_CONFIRMATION"`
	MailSuccessRedirect  string `json:"mail_success" env:"MAIL_SUCCESS"`
	MailFailedRedirect   string `json:"mail_failed" env:"MAIL_FAILED"`
	MailPassword         string `json:"mail_password" env:"MAIL_PASSWORD"`
	MailAddress          string `json:"mail_address" env:"MAIL_ADDRESS"`
	SMTPAddress          string `json:"mail_smtp" env:"MAIL_SMTP"`
	SMTPPort             string `json:"mail_smtp_port" env:"MAIL_SMTP_PORT"`

	AccessTokenValidityTime  time.Duration `json:"access_token_validity_time" env:"ACCESS_TOKEN_VALIDITY_TIME"`
	RefreshTokenValidityTime time.Duration `json:"refresh_token_validity_time" env:"REFRESH_TOKEN_VALIDITY_TIME"`
}
