package controllers

type ApiResponse struct {

	Code int32 `json:"code,omitempty"`

	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`
}
