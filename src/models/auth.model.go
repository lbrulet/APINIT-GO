package models

// LoginPayload is the login model
type LoginPayload struct {
	Username string `json:"username" binding:"required" example:"sankamille"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// RegisterPayload is the register model
type RegisterPayload struct {
	Username string `json:"username" binding:"required" example:"sankamille"`
	Email    string `json:"email" binding:"required" example:"luc.brulet@epitech.eu"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// RecoveryPayload is the recovery model
type RecoveryPayload struct {
	Email string `json:"email" binding:"required" exemple:"luc.brulet@epitech.eu"`
}
