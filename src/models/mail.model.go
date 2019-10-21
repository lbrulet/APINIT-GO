package models

// Template struct is used to fill to mail template
type Template struct {
	Email        string
	Username     string
	ConfirmEmail string
}

// TemplateRecovery struct is used to fill the mail recovery template
type TemplateRecovery struct {
	Email    string
	Username string
	Password string
}
