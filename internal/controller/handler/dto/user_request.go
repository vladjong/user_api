package dto

import "net/http"

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (u *CreateUserRequest) Bind(r *http.Request) error {
	return nil
}
