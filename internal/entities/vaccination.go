package entities

type Vaccination struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	DrugID string `json:"drug_id"`
	Dose   int    `json:"dose"`
	Date   string `json:"date"`
}
