package models

type Food struct {
	ID       int    `json:"id"`
	Name     string `json:"nama"`
	Category string `json:"category"`
	Price    string `json:"price"`
}
