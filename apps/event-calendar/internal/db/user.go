package db

import (
	"fmt"
	"github.com/harshabangi/LLD/apps/event-calendar/internal/models"
)

type userStorage struct {
	userMap map[string]models.User
}

type UserDB interface {
	CreateUser(info models.User)
	EnsureUser(name string) error
	GetUser(name string) (*models.User, error)
}

func NewUserDB() UserDB {
	return &userStorage{
		userMap: make(map[string]models.User),
	}
}

func (u userStorage) CreateUser(info models.User) {
	u.userMap[info.Name] = info
}

func (u userStorage) EnsureUser(name string) error {
	_, ok := u.userMap[name]
	if !ok {
		return fmt.Errorf("unknown user: %s", name)
	}
	return nil
}

func (u userStorage) GetUser(name string) (*models.User, error) {
	res, ok := u.userMap[name]
	if !ok {
		return nil, fmt.Errorf("unknown user: %s", name)
	}
	return &res, nil
}
