package user

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"test/internal/entities"
	"testing"
)

type executorMock struct {
	mock.Mock
}

func (m *executorMock) Registry(user entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *executorMock) Login(email, pass string) error {
	args := m.Called(email, pass)
	return args.Error(0)
}

func (m *executorMock) Exist(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUser_Login_Successful(t *testing.T) {
	storeMock := new(executorMock)
	UserServices := NewUser(storeMock)
	u := entities.User{
		Id:       "1234567",
		Name:     "TEST",
		Email:    "TEST@gmail.com",
		Password: "********",
	}
	storeMock.On("Login", u.Email, u.Password).Return(nil)
	err := UserServices.Login(u.Email, u.Password)
	assert.NoError(t, err)
}

func TestUser_Login_With_Error_When_Login_Return_nil(t *testing.T) {
	storeMock := new(executorMock)
	UserServices := NewUser(storeMock)
	u := entities.User{
		Id:       "1234567",
		Name:     "TEST",
		Email:    "TEST@gmail.com",
		Password: "********",
	}
	storeMock.On("Login", u.Email, u.Password).Return(errors.New("error"))
	err := UserServices.Login(u.Email, u.Password)
	assert.Error(t, err)
}

func TestUser_Exits_Successful(t *testing.T) {
	storeMock := new(executorMock)
	UserServices := NewUser(storeMock)
	u := entities.User{
		Id:       "1234567",
		Name:     "TEST",
		Email:    "TEST@gmail.com",
		Password: "********",
	}
	storeMock.On("Exist", u.Email).Return(nil)
	err := UserServices.Exits(u.Email)
	assert.NoError(t, err)
}

func TestUser_Exits_With_Error_When_Exist_Return_Error(t *testing.T) {
	storeMock := new(executorMock)
	UserServices := NewUser(storeMock)
	u := entities.User{
		Id:       "1234567",
		Name:     "TEST",
		Email:    "TEST@gmail.com",
		Password: "********",
	}
	storeMock.On("Exist", u.Email).Return(errors.New("error"))
	err := UserServices.Exits(u.Email)
	assert.Error(t, err)
}

func TestUser_Registry_Successful(t *testing.T) {
	storeMock := new(executorMock)
	UserServices := NewUser(storeMock)
	u := entities.User{
		Id:       "1234567",
		Name:     "TEST",
		Email:    "TEST@gmail.com",
		Password: "********",
	}
	storeMock.On("Exist", u.Email).Return(nil)
	storeMock.On("Registry", u).Return(nil)

	err := UserServices.Registry(u)
	assert.NoError(t, err)
}

func TestUser_Registry_With_Error_When_Exist_Return_Error(t *testing.T) {
	storeMock := new(executorMock)
	UserServices := NewUser(storeMock)
	u := entities.User{
		Id:       "1234567",
		Name:     "TEST",
		Email:    "TEST@gmail.com",
		Password: "********",
	}
	storeMock.On("Exist", u.Email).Return(errors.New("error"))
	storeMock.On("Registry", u).Return(nil)

	err := UserServices.Registry(u)
	assert.Error(t, err)
}
