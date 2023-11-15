package vaccination

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"test/internal/entities"
	"testing"
)

type executorDrugMock struct {
	mock.Mock
}

func (m *executorDrugMock) Registry(drug entities.Drug) error {
	args := m.Called(drug)
	return args.Error(0)
}

func (m *executorDrugMock) GetAllItems() ([]entities.Drug, error) {
	args := m.Called()
	return args.Get(0).([]entities.Drug), args.Error(1)
}

func (m *executorDrugMock) UpdateItem(id string, drug entities.Drug) error {
	args := m.Called(id, drug)
	return args.Error(0)
}

func (m *executorDrugMock) Exist(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *executorDrugMock) DeleteItem(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *executorDrugMock) GeItem(id string) (entities.Drug, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Drug), args.Error(1)
}

type executorVaccinationMock struct {
	mock.Mock
}

func (m *executorVaccinationMock) Registry(user entities.Vaccination) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *executorVaccinationMock) GetAllItems() ([]entities.Vaccination, error) {
	args := m.Called()
	return args.Get(0).([]entities.Vaccination), args.Error(1)
}

func (m *executorVaccinationMock) UpdateItem(id string, vaccination entities.Vaccination) error {
	args := m.Called(id, vaccination)
	return args.Error(0)
}

func (m *executorVaccinationMock) Exist(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *executorVaccinationMock) DeleteItem(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestVaccination_Registry_Successful(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := entities.Vaccination{
		ID:     "123",
		Name:   "TEST",
		DrugID: "6789",
		Dose:   10,
		Date:   "2023-11-15",
	}

	mockExecutorDrug.On("Exist", v.DrugID).Return(nil)
	mockExecutorDrug.On("GeItem", v.DrugID).Return(entities.Drug{
		ID:          "6789",
		Name:        "TEST",
		Approved:    false,
		MinDose:     5,
		MaxDose:     20,
		AvailableAt: "2023-12-25",
	}, nil)
	mockExecutorVaccination.On("Registry", v).Return(nil)
	err := service.Registry(v)

	assert.NoError(t, err)
}

func TestVaccination_Registry_With_Error_When_Exist_Return_Error(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := entities.Vaccination{
		ID:     "123",
		Name:   "TEST",
		DrugID: "6789",
		Dose:   10,
		Date:   "2023-11-15",
	}

	mockExecutorDrug.On("Exist", v.DrugID).Return(errors.New("error"))
	mockExecutorDrug.On("GeItem", v.DrugID).Return(entities.Drug{
		ID:          "6789",
		Name:        "TEST",
		Approved:    false,
		MinDose:     5,
		MaxDose:     20,
		AvailableAt: "2023-12-25",
	}, nil)
	mockExecutorVaccination.On("Registry", v).Return(nil)
	err := service.Registry(v)

	assert.Error(t, err)
}

func TestVaccination_Registry_With_Error_When_GeItem_Return_Error(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := entities.Vaccination{
		ID:     "123",
		Name:   "TEST",
		DrugID: "6789",
		Dose:   10,
		Date:   "2023-11-15",
	}

	mockExecutorDrug.On("Exist", v.DrugID).Return(nil)
	mockExecutorDrug.On("GeItem", v.DrugID).Return(entities.Drug{
		ID:          "6789",
		Name:        "TEST",
		Approved:    false,
		MinDose:     5,
		MaxDose:     20,
		AvailableAt: "2023-12-25",
	}, errors.New("error"))
	mockExecutorVaccination.On("Registry", v).Return(nil)
	err := service.Registry(v)

	assert.Error(t, err)
}

func TestVaccination_Registry_With_Error_When_Dose_Out_Of_Range(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := entities.Vaccination{
		ID:     "123",
		Name:   "TEST",
		DrugID: "6789",
		Dose:   1,
		Date:   "2023-11-15",
	}

	mockExecutorDrug.On("Exist", v.DrugID).Return(nil)
	mockExecutorDrug.On("GeItem", v.DrugID).Return(entities.Drug{
		ID:          "6789",
		Name:        "TEST",
		Approved:    false,
		MinDose:     5,
		MaxDose:     20,
		AvailableAt: "2023-12-25",
	}, nil)
	mockExecutorVaccination.On("Registry", v).Return(nil)
	err := service.Registry(v)

	assert.Error(t, err)
}

func TestVaccination_Registry_With_Error_When_Expired_Date_Vaccination(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := entities.Vaccination{
		ID:     "123",
		Name:   "TEST",
		DrugID: "6789",
		Dose:   10,
		Date:   "2023-11-15",
	}

	mockExecutorDrug.On("Exist", v.DrugID).Return(nil)
	mockExecutorDrug.On("GeItem", v.DrugID).Return(entities.Drug{
		ID:          "6789",
		Name:        "TEST",
		Approved:    false,
		MinDose:     5,
		MaxDose:     20,
		AvailableAt: "2023-10-25",
	}, nil)
	mockExecutorVaccination.On("Registry", v).Return(nil)
	err := service.Registry(v)

	assert.Error(t, err)
}

func TestVaccination_UpdateItem_Successful(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := entities.Vaccination{
		ID:     "123",
		Name:   "TEST",
		DrugID: "6789",
		Dose:   10,
		Date:   "2023-11-15",
	}

	mockExecutorVaccination.On("Exist", v.ID).Return(nil)
	mockExecutorVaccination.On("UpdateItem", v.ID, v).Return(nil)
	err := service.UpdateItem(v.ID, v)

	assert.NoError(t, err)
}

func TestVaccination_UpdateItem_With_Error_When_Exist_Return_Eror(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := entities.Vaccination{
		ID:     "123",
		Name:   "TEST",
		DrugID: "6789",
		Dose:   10,
		Date:   "2023-11-15",
	}

	mockExecutorVaccination.On("Exist", v.ID).Return(errors.New("error"))
	mockExecutorVaccination.On("UpdateItem", v.ID, v).Return(nil)
	err := service.UpdateItem(v.ID, v)

	assert.Error(t, err)
}

func TestVaccination_DeleteItem_Successful(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := entities.Vaccination{
		ID:     "123",
		Name:   "TEST",
		DrugID: "6789",
		Dose:   10,
		Date:   "2023-11-15",
	}

	mockExecutorVaccination.On("Exist", v.ID).Return(nil)
	mockExecutorVaccination.On("DeleteItem", v.ID).Return(nil)
	err := service.DeleteItem(v.ID)

	assert.NoError(t, err)
}

func TestVaccination_DeleteItem_With_Error_When_Exist_Return_Error(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := entities.Vaccination{
		ID:     "123",
		Name:   "TEST",
		DrugID: "6789",
		Dose:   10,
		Date:   "2023-11-15",
	}

	mockExecutorVaccination.On("Exist", v.ID).Return(errors.New("error"))
	mockExecutorVaccination.On("DeleteItem", v.ID).Return(nil)
	err := service.DeleteItem(v.ID)

	assert.Error(t, err)
}

func TestVaccination_GetAllItems_Successful(t *testing.T) {
	mockExecutorDrug := new(executorDrugMock)
	mockExecutorVaccination := new(executorVaccinationMock)
	service := NewVaccination(mockExecutorVaccination, mockExecutorDrug)
	v := []entities.Vaccination{
		{
			ID:     "123",
			Name:   "TEST",
			DrugID: "6789",
			Dose:   10,
			Date:   "2023-11-15",
		},
	}

	mockExecutorVaccination.On("GetAllItems").Return(v, nil)
	response, err := service.GetAllItems()

	assert.NoError(t, err)
	assert.Equal(t, response, v)
}
