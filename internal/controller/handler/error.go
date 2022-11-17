package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrIncorrectUserId = errors.New("incorrect user id")
)

type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	AppCode        int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error, statusCode int) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: statusCode,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
