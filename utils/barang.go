package utils

import (
	"kelompok3/toko-retail/config"
	"kelompok3/toko-retail/model"
	"time"
)

func CreateBarang(data model.Barang) (model.Barang, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetBarang() ([]model.Barang, error) {
	var barang model.Barang
	return barang.GetAll(config.Mysql.DB)
}

func GetBarangByID(id uint) (model.Barang, error) {
	barang := model.Barang{
		ID: id,
	}
	return barang.GetByID(config.Mysql.DB)
}

func UpdateBarang(id uint, barang model.Barang) (model.Barang, error) {
	barang.UpdatedAt = time.Now()
	err := barang.Create(config.Mysql.DB)

	return barang, err
}
