package utils

import (
	"kelompok3/toko-retail/config"
	"kelompok3/toko-retail/model"
	"time"
)

func CreateKodeDiskon(data model.Diskon) (model.Diskon, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetDiskon() ([]model.Diskon, error) {
	var Diskon model.Diskon
	return Diskon.GetAll(config.Mysql.DB)
}

func GetDiskonByCode(s string) (model.Diskon, error) {
	diskon := model.Diskon{
		KodeDiskon: s,
	}
	return diskon.GetByCode(config.Mysql.DB)
}

func GetDiskonByID(id uint) (model.Diskon, error) {
	Diskon := model.Diskon{
		ID: id,
	}
	return Diskon.GetByID(config.Mysql.DB)
}

func UpdateDiskon(id uint, Diskon model.Diskon) (model.Diskon, error) {
	Diskon.UpdatedAt = time.Now()
	err := Diskon.Create(config.Mysql.DB)

	return Diskon, err
}
