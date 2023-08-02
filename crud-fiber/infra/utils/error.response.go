package utils

type (
	IError struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	ErrorResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Error   string `json:"err"`
	}
)
