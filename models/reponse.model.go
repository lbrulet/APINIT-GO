package models

// ResponsePayload is the api response
type ResponsePayload struct {
	Success bool   `json:"success" binding:"required"`
	Message string `json:"message" binding:"required"`
}
