package model

import "gorm.io/gorm"

type Barang struct {
	ID         uint64  `gorm:"primarykey" json:"id"`
	KodeBarang string  `json:"kode_barang"`
	Nama       string  `json:"nama_barang"`
	HargaPokok float64 `json:"harga_pokok"`
	HargaJual  float64 `json:"harga_jual"`
	TipeBarang string  `json:"tipe_barang"`
	Stok       uint    `json:"stok"`
	Model
	CreatedBy string `json:"created_by"`
}

type Details struct {
	ID         uint64  `gorm:"primarykey" json:"id"`
	KodeBarang string  `json:"kode_barang"`
	Nama       string  `json:"nama_barang"`
	HargaPokok float64 `json:"harga_pokok"`
	HargaJual  float64 `json:"harga_jual"`
	TipeBarang string  `json:"tipe_barang"`
	Stok       uint    `json:"stok"`
	Model
	CreatedBy string        `json:"created_by"`
	Histori   []HistoriASKM `gorm:"foreignKey:ID_Barang" json:"histori_stok"`
}

type CreateB struct {
	ID         uint64       `gorm:"primarykey" json:"id"`
	KodeBarang string       `json:"kode_barang"`
	Nama       string       `json:"nama_barang"`
	HargaPokok float64      `json:"harga_pokok"`
	HargaJual  float64      `json:"harga_jual"`
	TipeBarang string       `json:"tipe_barang"`
	Stok       uint         `json:"stok"`
	CreatedBy  string       `json:"created_by"`
	Histori    []HistoriASK `gorm:"foreignKey:ID_Barang" json:"histori_stok"`
}

func (br *Barang) Create(db *gorm.DB) error {
	err := db.
		Model(Barang{}).
		Create(&br).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (br *Barang) GetAll(db *gorm.DB) ([]Barang, error) {
	res := []Barang{}

	err := db.
		Model(Barang{}).Order("created_at desc").
		Limit(50).
		Find(&res).
		Error

	if err != nil {
		return []Barang{}, err
	}

	return res, nil
}

func (br *Barang) GetByID(db *gorm.DB) (Barang, error) {
	res := Barang{}

	err := db.
		Model(Barang{}).
		Where("id = ?", br.ID).
		Take(&res).
		Error

	if err != nil {
		return Barang{}, err
	}

	return res, nil
}

func (br *Barang) Update(db *gorm.DB) error {
	err := db.
		Model(&Barang{}).
		Where("id = ?", br.ID).
		Updates(&br).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (br *Barang) Delete(db *gorm.DB) error {
	err := db.
		Model(&Barang{}).
		Where("id = ?", br.ID).
		Delete(&br).
		Error

	if err != nil {
		return err
	}

	return nil
}
