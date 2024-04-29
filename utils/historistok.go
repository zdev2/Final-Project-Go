package utils

import (
	"kelompok3/toko-retail/config"
	"kelompok3/toko-retail/model"
)

func CreateHistoriBarang(p *model.Details, keterangan string, amount int, status string) (model.Histori, error) {
	Histori := model.Histori{
		ID_barang:  uint(p.ID),
		Amount:     amount,
		Status:     status,
		Keterangan: keterangan,
	}

	return Histori, Histori.Create(config.Mysql.DB)
}

func CreateHistoriPenjualan(p *model.CreateP, keterangan string, amount int, status string) (model.Histori, error) {
	Histori := model.Histori{
		ID_barang:  uint(p.ID),
		Amount:     amount,
		Status:     status,
		Keterangan: keterangan,
	}

	return Histori, Histori.Create(config.Mysql.DB)
}

func GetASKMByIDBarang(idb uint64) ([]model.HistoriASKM, error) {
	histori := model.Histori{
		ID_barang: uint(idb),
	}

	newHistory, err := histori.GetIDBarang(config.Mysql.DB)
	if err != nil {
		return nil, err
	}

	var haskm []model.HistoriASKM
	for _, h := range newHistory {
		haskm = append(haskm, model.HistoriASKM{
			Amount:     h.Amount,
			Status:     h.Status,
			Keterangan: h.Keterangan,
			Model:      h.Model,
		})
	}

	return haskm, err
}

func GetASK(idb uint64) ([]model.HistoriASK, error) {
	histori := model.Histori{
		ID_barang: uint(idb),
	}

	newHistory, err := histori.GetIDBarang(config.Mysql.DB)
	if err != nil {
		return nil, err
	}

	var hask []model.HistoriASK
	for _, h := range newHistory {
		hask = append(hask, model.HistoriASK{
			Amount:     h.Amount,
			Status:     h.Status,
			Keterangan: h.Keterangan,
		})
	}

	return hask, err
}
