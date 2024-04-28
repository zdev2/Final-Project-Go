package utils

import (
	"kelompok3/toko-retail/config"
	"kelompok3/toko-retail/model"
	"time"
)

func InsertPenjualanData(data model.Penjualan) (model.Penjualan, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.CreatePenjualan(config.Mysql.DB)

	return data, err
}

func GetPenjualan() ([]model.Penjualan, error) {
	var penjualan model.Penjualan
	return penjualan.GetAll(config.Mysql.DB)
}

// func GetPenjualanByID(id uint) (model.Penjualan, error) {
// 	penjualan := model.Penjualan{
// 		ID: id,
// 	}
// 	return penjualan.GetByID(config.Mysql.DB)
// }
