package vaccination

import (
	"errors"
	"test/internal/entities"
	"test/internal/repository/drug"
	vaccinationRepository "test/internal/repository/vaccination"
)

type Executor interface {
	Registry(vaccination entities.Vaccination) error
	UpdateItem(id string, vaccination entities.Vaccination) error
	GetAllItems() ([]entities.Vaccination, error)
	DeleteItem(id string) error
	Exist(id string) error
}

type vaccination struct {
	vaccinationStorage vaccinationRepository.Executor
	drugStorage        drug.Executor
}

func NewVaccination(vaccinationStorage vaccinationRepository.Executor, drugStorage drug.Executor) Executor {
	return &vaccination{
		vaccinationStorage: vaccinationStorage,
		drugStorage:        drugStorage,
	}
}

func (v *vaccination) Registry(vaccination entities.Vaccination) error {
	err := v.drugStorage.Exist(vaccination.DrugID)
	if err != nil {
		return err
	}

	item, err := v.drugStorage.GeItem(vaccination.DrugID)
	if err != nil {
		return err
	}

	if vaccination.Dose < item.MinDose || vaccination.Dose > item.MaxDose {
		return errors.New("dose out of permitted range")
	}

	if vaccination.Date >= item.AvailableAt {
		return errors.New("expired vaccination date")
	}

	return v.vaccinationStorage.Registry(vaccination)
}

func (v *vaccination) UpdateItem(id string, vaccination entities.Vaccination) error {
	err := v.Exist(id)
	if err != nil {
		return err
	}
	return v.vaccinationStorage.UpdateItem(id, vaccination)
}

func (v *vaccination) GetAllItems() ([]entities.Vaccination, error) {
	return v.vaccinationStorage.GetAllItems()
}

func (v *vaccination) DeleteItem(id string) error {
	err := v.Exist(id)
	if err != nil {
		return err
	}

	return v.vaccinationStorage.DeleteItem(id)
}

func (v *vaccination) Exist(id string) error {
	return v.vaccinationStorage.Exist(id)
}
