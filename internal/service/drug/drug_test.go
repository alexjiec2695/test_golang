package drug

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

func (m *executorMock) Registry(drug entities.Drug) error {
	args := m.Called(drug)
	return args.Error(0)
}

func (m *executorMock) GetAllItems() ([]entities.Drug, error) {
	args := m.Called()
	return args.Get(0).([]entities.Drug), args.Error(1)
}

func (m *executorMock) UpdateItem(id string, drug entities.Drug) error {
	args := m.Called(id, drug)
	return args.Error(0)
}

func (m *executorMock) Exist(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *executorMock) DeleteItem(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *executorMock) GeItem(id string) (entities.Drug, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Drug), args.Error(1)
}

func TestDrugs_Registry_Successful(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)
	d := entities.Drug{
		ID:          "123456",
		Name:        "TEST",
		Approved:    false,
		MinDose:     900,
		MaxDose:     1200,
		AvailableAt: "2013-11-15",
	}
	storeMock.On("Registry", d).Return(nil)
	err := drugService.Registry(d)
	assert.NoError(t, err)
}

func TestDrugs_Registry_With_Error_When_Registry_Return_Error(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)
	d := entities.Drug{
		ID:          "123456",
		Name:        "TEST",
		Approved:    false,
		MinDose:     900,
		MaxDose:     1200,
		AvailableAt: "2013-11-15",
	}
	storeMock.On("Registry", d).Return(errors.New("error"))
	err := drugService.Registry(d)
	assert.Error(t, err)
	assert.EqualError(t, err, "error")
}

func TestDrugs_UpdateItem_Successful(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)
	d := entities.Drug{
		ID:          "123456",
		Name:        "TEST",
		Approved:    false,
		MinDose:     900,
		MaxDose:     1200,
		AvailableAt: "2013-11-15",
	}
	storeMock.On("Exist", "123456").Return(nil)
	storeMock.On("UpdateItem", "123456", d).Return(nil)
	err := drugService.UpdateItem("123456", d)
	assert.NoError(t, err)
}

func TestDrugs_UpdateItem_With_Error_When_Exist_Return_Error(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)
	d := entities.Drug{
		ID:          "123456",
		Name:        "TEST",
		Approved:    false,
		MinDose:     900,
		MaxDose:     1200,
		AvailableAt: "2013-11-15",
	}
	storeMock.On("Exist", "123456").Return(errors.New("error"))
	err := drugService.UpdateItem("123456", d)
	assert.Error(t, err)
	assert.EqualError(t, err, "error")
}

func TestDrugs_UpdateItem_With_Error_When_UpdateItem_Return_Error(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)
	d := entities.Drug{
		ID:          "123456",
		Name:        "TEST",
		Approved:    false,
		MinDose:     900,
		MaxDose:     1200,
		AvailableAt: "2013-11-15",
	}
	storeMock.On("Exist", "123456").Return(nil)
	storeMock.On("UpdateItem", "123456", d).Return(errors.New("error"))
	err := drugService.UpdateItem("123456", d)
	assert.Error(t, err)
	assert.EqualError(t, err, "error")
}

func TestDrugs_GetAllItems_Successful(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)
	d := []entities.Drug{
		{
			ID:          "123456",
			Name:        "TEST",
			Approved:    false,
			MinDose:     900,
			MaxDose:     1200,
			AvailableAt: "2013-11-15",
		},
	}
	storeMock.On("GetAllItems").Return(d, nil)
	result, err := drugService.GetAllItems()
	assert.NoError(t, err)
	assert.Equal(t, result, d)
}

func TestDrugs_GetAllItems_With_Error_When_GetAllItem_Return_Error(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)
	d := []entities.Drug{
		{
			ID:          "123456",
			Name:        "TEST",
			Approved:    false,
			MinDose:     900,
			MaxDose:     1200,
			AvailableAt: "2013-11-15",
		},
	}
	storeMock.On("GetAllItems").Return(d, errors.New("error"))
	_, err := drugService.GetAllItems()
	assert.Error(t, err)
	assert.EqualError(t, err, "error")
}

func TestDrugs_DeleteItem_Successful(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)

	storeMock.On("Exist", "123456").Return(nil)
	storeMock.On("DeleteItem", "123456").Return(nil)
	err := drugService.DeleteItem("123456")
	assert.NoError(t, err)
}

func TestDrugs_DeleteItem_With_Error_When_Delete_Return_Error(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)

	storeMock.On("Exist", "123456").Return(errors.New("error"))
	storeMock.On("DeleteItem", "123456").Return(nil)
	err := drugService.DeleteItem("123456")
	assert.Error(t, err)
	assert.EqualError(t, err, "error")
}

func TestDrugs_Exist_Successful(t *testing.T) {
	storeMock := new(executorMock)
	drugService := NewDrug(storeMock)

	storeMock.On("Exist", "123456").Return(nil)
	err := drugService.Exist("123456")
	assert.NoError(t, err)
}
