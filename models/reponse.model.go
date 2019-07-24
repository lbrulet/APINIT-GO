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
	RefreshToken string `json:"refresh-token" binding:"required"`
}
