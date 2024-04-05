package model

type Diskon struct {
  Kode_diskon string `json:"kode_diskon"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Model
}
