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

func DeleteBarang(id uint64) error {
	barang := model.Barang{
		ID: id,
	}
	return barang.Delete(config.Mysql.DB)
}

// {
// 	"nama_barang":"nasi",
// 	"harga_pokok":3000,
// 	"harga_jual":5000,
// 	"tipe_barang":"MAKANAN",
// 	"stok":20
//   }

// "data": {
//     "id": 1,
//     "kode_invoice": "INV/1",
//     "nama pembeli": "asep",
//     "subtotal": 20000,
//     "kode_diskon": null,
//     "diskon": 0,
//     "total": 20000,
//     "created_at": "2023-10-31 00:00:00",
//     "updated_at": "2023-10-31 00:00:00",
//     "deleted_at": null,
//     "created_by": "SYSTEM",
//     "item_penjualan": [
//       {
//         "kode_barang": "MI-1",
//         "jumlah": 1,
//         "subtotal": 10000,
//         "created_at": "2023-10-31 00:00:00",
//         "updated_at": "2023-10-31 00:00:00",
//         "deleted_at": null
//       },
