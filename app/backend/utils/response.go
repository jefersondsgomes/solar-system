package utils

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func GenerateErrorResponse(statusCode int, message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: statusCode,
		Error:      message,
	}
}
