package errors

import "net/http"

// RestErr error struct
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequest error message
func NewBadRequest(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFound error message
func NewNotFound(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  http.StatusNotFound,
		Error:   "not found",
	}
}
