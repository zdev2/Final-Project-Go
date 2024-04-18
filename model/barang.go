package model

type Barang struct {
	KodeBarang string  `json:"kode_barang"`
	Nama       string  `json:"nama"`
	HargaPokok float64 `json:"harga_pokok"`
	HargaJual  float64 `json:"harga_jual"`
	TipeBarang string  `json:"tipe_barang"`
	Stok       uint    `json:"stok"`
	Model
	CreatedBy string `json:"created_by"`
}
