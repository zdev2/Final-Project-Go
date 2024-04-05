package model

type Histori struct {
  ID_barang uint `json:"id_barang"`
  Amount int `json:"amount"`
  Status string `json:"status"`
  Keterangan string`json:"keterangan"`
  Model
}