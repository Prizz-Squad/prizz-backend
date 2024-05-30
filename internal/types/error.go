package types

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var apiError Error
	if errors.As(err, &apiError) {
		return c.Status(apiError.Code).JSON(apiError)
	}
	apiError = NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

func (e Error) Error() string {
	return e.Err
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func ErrUnAuthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err:  "unauthorized request",
	}
}

func ErrNotResourceNotFound(res string) Error {
	return Error{
		Code: http.StatusNotFound,
		Err:  res + " resource not found",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid JSON request",
	}
}

func ErrInvalidID() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid id given",
	}
}

func InternalServerError() Error {
	return Error{
		Code: http.StatusInternalServerError,
		Err:  "Internal server error",
	}
}
