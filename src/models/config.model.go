package models

import "time"

// Configuration struct is used to load configuration from file or environnement
type Configuration struct {
	Port                       string `json:"app_port" env:"APP_PORT"`
	DatabaseName               string `json:"database_name" env:"DATABASE_NAME"`
	DatabaseEndPoint           string `json:"database_endpoint" env:"DATABASE_ENDPOINT"`
	DatabaseUser               string `json:"database_user" env:"DATABASE_USER"`
	DatabasePassword           string `json:"database_password" env:"DATABASE_PASSWORD"`
	MailFrom                   string `json:"mail_from" env:"MAIL_FROM"`
	MailSubjectConfirmAccount  string `json:"mail_subject_confirm_account" env:"MAIL_SUBJECT_CONFIRM_ACCOUNT"`
	MailSubjectRecoveryAccount string `json:"mail_subject_recovery_account" env:"MAIL_SUBJECT_RECOVERY_ACCOUNT"`
	MailConfirmationLink       string `json:"mail_confirmation" env:"MAIL_CONFIRMATION"`
	MailSuccessRedirect        string `json:"mail_success" env:"MAIL_SUCCESS"`
	MailFailedRedirect         string `json:"mail_failed" env:"MAIL_FAILED"`
	MailPassword               string `json:"mail_password" env:"MAIL_PASSWORD"`
	MailAddress                string `json:"mail_address" env:"MAIL_ADDRESS"`
	SMTPAddress                string `json:"mail_smtp" env:"MAIL_SMTP"`
	SMTPPort                   string `json:"mail_smtp_port" env:"MAIL_SMTP_PORT"`

	AccessTokenValidityTime  time.Duration `json:"access_token_validity_time" env:"ACCESS_TOKEN_VALIDITY_TIME"`
	RefreshTokenValidityTime time.Duration `json:"refresh_token_validity_time" env:"REFRESH_TOKEN_VALIDITY_TIME"`
}
