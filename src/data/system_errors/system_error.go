package system_errors

import "net/http"

type SystemError struct {
	Message string
	Code    http.ConnState
}

func (e *SystemError) Error() string {
	return e.Message
}

func New(message string, code http.ConnState) error {
	return &SystemError{message, code}
}
