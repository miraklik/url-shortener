package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOk    = "OK"
	StatusError = "ERROR"
)

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("filed %s is a required field", err.Field()))
		case "url":
			errMsg = append(errMsg, fmt.Sprintf("filed %s is not a valid URL", err.Field()))
		default:
			errMsg = append(errMsg, fmt.Sprintf("filed %s is not valid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsg, ", "),
	}
}
