package apperror

import "net/http"

type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code, msg string, status int) *AppError {
	return &AppError{
		Code:       code,
		Message:    msg,
		HTTPStatus: status,
	}
}

var (
	ErrBadRequest = New(
		"REQ_400",
		"invalid request",
		http.StatusBadRequest,
	)

	ErrUnauthorized = New(
		"AUTH_401",
		"unauthorized",
		http.StatusUnauthorized,
	)

	ErrForbidden = New(
		"AUTH_403",
		"forbidden",
		http.StatusForbidden,
	)

	ErrNotFound = New(
		"GEN_404",
		"data not found",
		http.StatusNotFound,
	)

	ErrInternal = New(
		"GEN_500",
		"internal server error",
		http.StatusInternalServerError,
	)
)
