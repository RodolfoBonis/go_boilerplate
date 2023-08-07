package utils

import "net/http"

type HttpError struct {
	StatusCode int    `json:"code"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

func UnauthorizedError() HttpError {
	return HttpError{
		http.StatusUnauthorized,
		"Unauthorized",
		"You are not authorized to access this resource",
	}
}

func NewHTTPError(statusCode int, error string, message string) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Error:      error,
		Message:    message,
	}
}

func NotFoundError() *HttpError {
	return &HttpError{
		http.StatusNotFound,
		"Not found",
		"The requested resource was not found",
	}
}

func DataLayerError(message string) *HttpError {
	return &HttpError{
		http.StatusInternalServerError,
		"Data Layer Error",
		message,
	}
}

func RepositoryLayerError(message string) *HttpError {
	return &HttpError{
		http.StatusInternalServerError,
		"Repository Layer Error",
		message,
	}
}

func DtoLayerError(message string) *HttpError {
	return &HttpError{
		http.StatusInternalServerError,
		"DTO Layer Error",
		message,
	}
}

func InternalServerError(message string) *HttpError {
	return &HttpError{
		http.StatusInternalServerError,
		"Internal Server Error",
		message,
	}
}

func BadRequestError(message string) *HttpError {
	return &HttpError{
		http.StatusBadRequest,
		"Bad Request",
		message,
	}
}

func TokenExpiredError() *HttpError {
	return &HttpError{
		http.StatusUnauthorized,
		"Token Expired!",
		"Your token has expired, You are not authorized to access this resource!",
	}
}

func ForbiddenError(message string) *HttpError {
	return &HttpError{
		http.StatusForbidden,
		"Forbidden",
		message,
	}
}
