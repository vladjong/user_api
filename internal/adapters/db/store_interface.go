package db

import (
	"refactoring/internal/controller/handler/dto"
	"refactoring/internal/entities"
)

type UserStore interface {
	SearchUsers() (userStore entities.UserStore, err error)
	CreateUser(user entities.User) (id string, err error)
	GetUser(id string) (user entities.User, err error)
	UpdateUser(id string, userUpdate dto.UpdateUserRequest) error
	DeleteUser(id string) error
}
