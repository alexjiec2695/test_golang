package drug

import (
	"test/internal/entities"
	"test/internal/repository/drug"
)

type Executor interface {
	Registry(drug entities.Drug) error
	UpdateItem(id string, drug entities.Drug) error
	GetAllItems() ([]entities.Drug, error)
	DeleteItem(id string) error
	Exist(id string) error
}

type drugs struct {
	storage drug.Executor
}

func NewDrug(storage drug.Executor) Executor {
	return &drugs{
		storage: storage,
	}
}

func (d *drugs) Registry(drug entities.Drug) error {
	return d.storage.Registry(drug)
}

func (d *drugs) UpdateItem(id string, drug entities.Drug) error {
	err := d.Exist(id)
	if err != nil {
		return err
	}
	return d.storage.UpdateItem(id, drug)
}

func (d *drugs) GetAllItems() ([]entities.Drug, error) {
	return d.storage.GetAllItems()
}

func (d *drugs) DeleteItem(id string) error {
	err := d.Exist(id)
	if err != nil {
		return err
	}

	return d.storage.DeleteItem(id)
}

func (d *drugs) Exist(id string) error {
	return d.storage.Exist(id)
}
