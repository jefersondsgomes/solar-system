package helpers

type Empty struct{}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func GenerateErrorResponse(statusCode int, message string) ErrorResponse {
	response := ErrorResponse{
		StatusCode: statusCode,
		Error:      message,
	}

	return response
}
