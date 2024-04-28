package utils

import "kelompok3/toko-retail/model"

func CreateHistory(p *model.Details, keterangan string, amount int, status string) (*model.HistoriASKM, error) {
	history := model.Histori{
		ID_barang:  uint(p.ID),
		Amount:     amount,
		Status:     status,
		Keterangan: keterangan,
	}

	HistoriASKM := model.HistoriASKM{
		Amount:     history.Amount,
		Status:     history.Status,
		Keterangan: keterangan,
	}
	p.Histori = append(p.Histori, HistoriASKM)
}
