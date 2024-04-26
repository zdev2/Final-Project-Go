package utils

import (
	"fmt"
	"kelompok3/toko-retail/config"
	"kelompok3/toko-retail/model"
	"strconv"
	"time"
)

func CreateBarang(data model.Barang) (model.Barang, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)
	if data.TipeBarang == "MAKANAN" {
		data.KodeBarang = fmt.Sprintf("MA-%v", strconv.FormatUint(uint64(data.ID), 10))
	} else if data.TipeBarang == "MINUMAN" {
		data.KodeBarang = fmt.Sprintf("MI-%v", strconv.FormatUint(uint64(data.ID), 10))
	} else {
		data.KodeBarang = fmt.Sprintf("L-%v", strconv.FormatUint(uint64(data.ID), 10))
	}
	data.Update(config.Mysql.DB)

	return data, err
}

func GetBarang() ([]model.Barang, error) {
	var barang model.Barang
	return barang.GetAll(config.Mysql.DB)
}

func GetBarangByID(id uint64) (model.Details, error) {
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

// {
// 	"nama_barang":"nasi",
// 	"harga_pokok":3000,
// 	"harga_jual":5000,
// 	"tipe_barang":"MAKANAN",
// 	"stok":20
//   }
