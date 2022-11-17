package jsondb

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"refactoring/internal/controller/handler"
	"refactoring/internal/controller/handler/dto"
	"refactoring/internal/entities"
	"strconv"
)

const store = `users.json`

type storeStorage struct{}

func New() *storeStorage {
	return &storeStorage{}
}

func (s *storeStorage) SearchUsers() (userStore entities.UserStore, err error) {
	fileDate, err := os.ReadFile(store)
	if err != nil {
		return userStore, err
	}
	if err = json.Unmarshal(fileDate, &userStore); err != nil {
		return userStore, err
	}
	return userStore, err
}

func (s *storeStorage) CreateUser(user entities.User) (id string, err error) {
	userStore, err := s.SearchUsers()
	if err != nil {
		return "", err
	}
	userStore.Increment++
	id = strconv.Itoa(userStore.Increment)
	userStore.List[id] = user
	biteUserStore, err := json.Marshal(&userStore)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(store, biteUserStore, fs.ModePerm)
	return id, err
}

func (s *storeStorage) GetUser(id string) (user entities.User, err error) {
	userStore, err := s.SearchUsers()
	if err != nil {
		return user, err
	}
	user, ok := userStore.List[id]
	if !ok {
		return user, handler.ErrUserNotFound
	}
	return user, nil
}

func (s *storeStorage) UpdateUser(id string, userUpdate dto.UpdateUserRequest) error {
	userStore, err := s.SearchUsers()
	if err != nil {
		return err
	}
	user, ok := userStore.List[id]
	if !ok {
		return handler.ErrUserNotFound
	}
	user.DisplayName = userUpdate.DisplayName
	userStore.List[id] = user
	biteUserStore, _ := json.Marshal(&userStore)
	err = ioutil.WriteFile(store, biteUserStore, fs.ModePerm)
	return err
}

func (s *storeStorage) DeleteUser(id string) error {
	userStore, err := s.SearchUsers()
	if err != nil {
		return err
	}
	_, ok := userStore.List[id]
	if !ok {
		return handler.ErrUserNotFound
	}
	delete(userStore.List, id)
	userStore.Increment--
	biteUserStore, _ := json.Marshal(&userStore)
	err = ioutil.WriteFile(store, biteUserStore, fs.ModePerm)
	return err
}
