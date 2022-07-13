package utils

type ErrorResponse struct {
	err     string
	message string
}

func FormatErrorResponse(err string, message string) ErrorResponse {
	return ErrorResponse{err: err, message: message}
}
