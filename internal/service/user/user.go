package user

import (
	"test/internal/entities"
	userRepository "test/internal/repository/user"
)

type Executor interface {
	Registry(user entities.User) error
	Login(email, pass string) error
	Exits(email string) error
}

type user struct {
	storage userRepository.Executor
}

func NewUser(storage userRepository.Executor) Executor {
	return &user{
		storage: storage,
	}
}

func (s *user) Registry(user entities.User) error {
	err := s.Exits(user.Email)
	if err != nil {
		return err
	}

	return s.storage.Registry(user)
}

func (s *user) Login(email, pass string) error {
	return s.storage.Login(email, pass)
}

func (s *user) Exits(email string) error {
	return s.storage.Exist(email)
}
