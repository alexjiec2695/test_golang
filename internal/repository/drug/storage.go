package drug

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test/internal/entities"
	"test/internal/repository/postgres/data"
	"time"
)

type Executor interface {
	Registry(user entities.Drug) error
	GetAllItems() ([]entities.Drug, error)
	UpdateItem(id string, drug entities.Drug) error
	Exist(id string) error
	DeleteItem(id string) error
	GeItem(id string) (entities.Drug, error)
}

type storage struct {
	db *gorm.DB
}

func NewDrugRepository(db *gorm.DB) Executor {
	return &storage{
		db: db,
	}
}

func (s *storage) Registry(drug entities.Drug) error {
	available, err := time.Parse("2006-01-02", drug.AvailableAt)
	if err != nil {
		return err
	}

	drugData := data.Drug{
		ID:          uuid.New().String(),
		Name:        drug.Name,
		Approved:    drug.Approved,
		MinDose:     drug.MinDose,
		MaxDose:     drug.MaxDose,
		AvailableAt: available,
	}

	err = s.db.Save(drugData).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) GetAllItems() ([]entities.Drug, error) {
	var result []data.Drug
	tx := s.db.Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	resolve := make([]entities.Drug, len(result))
	for i, drug := range result {
		resolve[i] = entities.Drug{
			ID:          drug.ID,
			Name:        drug.Name,
			Approved:    drug.Approved,
			MinDose:     drug.MinDose,
			MaxDose:     drug.MaxDose,
			AvailableAt: drug.AvailableAt.String(),
		}
	}

	return resolve, nil

}

func (s *storage) GeItem(id string) (entities.Drug, error) {
	var result data.Drug
	tx := s.db.Where("id = ?", id).Find(&result)
	if tx.Error != nil {
		return entities.Drug{}, tx.Error
	}

	return entities.Drug{
		ID:          result.ID,
		Name:        result.Name,
		Approved:    result.Approved,
		MinDose:     result.MinDose,
		MaxDose:     result.MaxDose,
		AvailableAt: result.AvailableAt.Format("2006-01-02"),
	}, nil
}

func (s *storage) UpdateItem(id string, drug entities.Drug) error {
	result := s.db.Model(&data.Drug{}).Where("id = ?", id).Updates(drug)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *storage) Exist(id string) error {
	var result []data.Drug
	tx := s.db.Where("id = ?", id).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if len(result) == 0 {
		return errors.New("drug don't exist")
	}

	return nil
}

func (s *storage) DeleteItem(id string) error {
	result := s.db.Model(&data.Drug{}).Where("id = ?", id).Delete(&data.Drug{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
