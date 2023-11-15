package entities

type Drug struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Approved    bool   `json:"approved"`
	MinDose     int    `json:"min_dose"`
	MaxDose     int    `json:"max_dose"`
	AvailableAt string `json:"available_at"`
}
