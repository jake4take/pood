package models

type Unit struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	FirstForm   string `json:"first_form"`
	SecondForm  string `json:"second_form"`
	ThirdForm   string `json:"third_form"`
}
