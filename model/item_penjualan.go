package model

type ItemPenjualan struct {
	ID          uint    `json:"id"`
	IDBarang    uint    `json:"id_barang"`
	IDPenjualan uint    `json:"id_penjualan"`
	Jumlah      uint    `json:"jumlah"`
	SubTotal    float64 `json:"sub_total"`
	Model
}
