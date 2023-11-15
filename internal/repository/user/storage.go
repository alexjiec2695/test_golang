package user

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test/internal/entities"
	"test/internal/repository/postgres/data"
)

type Executor interface {
	Registry(user entities.User) error
	Login(email, pass string) error
	Exist(email string) error
}

type storage struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Executor {
	return &storage{
		db: db,
	}
}

func (s *storage) Registry(user entities.User) error {
	userData := data.User{
		ID:       uuid.New().String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	err := s.db.Save(userData).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) Login(email, pass string) error {
	var result []data.User
	tx := s.db.Where("email = ? and password = ?", email, pass).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if len(result) != 1 {
		return errors.New("email or password invalid")
	}

	return nil
}

func (s *storage) Exist(email string) error {
	var result []data.User
	tx := s.db.Where("email = ?", email).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if len(result) != 0 {
		return errors.New("email already exists")
	}
	return nil
}
