package vaccination

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test/internal/entities"
	"test/internal/repository/postgres/data"
)

type Executor interface {
	Registry(user entities.Vaccination) error
	GetAllItems() ([]entities.Vaccination, error)
	UpdateItem(id string, vaccination entities.Vaccination) error
	Exist(id string) error
	DeleteItem(id string) error
}

type storage struct {
	db *gorm.DB
}

func NewVaccinationRepository(db *gorm.DB) Executor {
	return &storage{
		db: db,
	}
}

func (s *storage) Registry(vaccination entities.Vaccination) error {
	vaccinationData := data.Vaccination{
		ID:     uuid.New().String(),
		Name:   vaccination.Name,
		DrugID: vaccination.DrugID,
		Dose:   vaccination.Dose,
		Date:   vaccination.Date,
	}

	err := s.db.Save(vaccinationData).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) GetAllItems() ([]entities.Vaccination, error) {
	var result []data.Vaccination
	tx := s.db.Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	resolve := make([]entities.Vaccination, len(result))
	for i, vaccination := range result {
		resolve[i] = entities.Vaccination{
			ID:     vaccination.ID,
			Name:   vaccination.Name,
			DrugID: vaccination.DrugID,
			Dose:   vaccination.Dose,
			Date:   vaccination.Date,
		}
	}

	return resolve, nil

}

func (s *storage) UpdateItem(id string, vaccination entities.Vaccination) error {
	result := s.db.Model(&data.Vaccination{}).Where("id = ?", id).Updates(vaccination)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *storage) Exist(id string) error {
	var result []data.Vaccination
	tx := s.db.Where("id = ?", id).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if len(result) == 0 {
		return errors.New("vaccination don't exist")
	}

	return nil
}

func (s *storage) DeleteItem(id string) error {
	result := s.db.Model(&data.Vaccination{}).Where("id = ?", id).Delete(&data.Vaccination{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
