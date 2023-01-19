package models

import "encoding/json"

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Rate     float64 `json:"rate"`
	Amount   float64 `json:"amount"`
	Raised   float64 `json:"raised"`
	Cycle    int     `json:"cycle"`
	Deadline string  `json:"deadline"`
}

type ProductSlice []Product

func (*Product) TableName() string {
	return "product"
}

func (p Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p ProductSlice) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *ProductSlice) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}
