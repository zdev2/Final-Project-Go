package model

import "gorm.io/gorm"

type Barang struct {
	ID         uint    `gorm:"primarykey" json:"id"`
	KodeBarang string  `json:"kode_barang"`
	Nama       string  `json:"nama"`
	HargaPokok float64 `json:"harga_pokok"`
	HargaJual  float64 `json:"harga_jual"`
	TipeBarang string  `json:"tipe_barang"`
	Stok       uint    `json:"stok"`
	Model
	CreatedBy string `json:"created_by"`
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
		Model(Barang{}).
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
