package data

type Vaccination struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	DrugID string `json:"drug_id"`
	Dose   int    `json:"dose"`
	Date   string `json:"date"`
}
