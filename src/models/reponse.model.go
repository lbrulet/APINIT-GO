package models

// ResponsePayload is the api response
type ResponsePayload struct {
	Success bool        `json:"success" binding:"required"`
	Message interface{} `json:"message" binding:"required"`
}

// LoginResponsePayload is the login response
type LoginResponsePayload struct {
	Success      bool   `json:"success" binding:"required"`
	Message      string `json:"message" binding:"required"`
	Token        string `json:"token" binding:"required"`
	User         User   `json:"user" binding:"required"`
	RefreshToken string `json:"refresh-token" binding:"required"`
}

// RegisterResponsePayload is the register response
type RegisterResponsePayload struct {
	Success bool   `json:"success" binding:"required"`
	Message string `json:"message" binding:"required"`
	Token   string `json:"token" binding:"required"`
	User    User   `json:"user" binding:"required"`
}
