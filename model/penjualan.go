package model

type Penjualan struct {
	ID           uint    `gorm:"primarykey" json:"id"`
	Kode_invoice string  `json:"kode_invoice"`
	Nama_pembeli string  `json:"nama_pembeli"`
	Subtotal     float64 `json:"subtotal"`
	Kode_diskon  string  `json:"kode_diskon"`
	Diskon       float64 `json:"diskon"`
	Total        float64 `json:"total"`
	Model
	Created_by string `json:"created_by"`
}
