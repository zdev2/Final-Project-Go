package controllers

import (
	"kelompok3/toko-retail/model"
	"kelompok3/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

/*
TO DO:

	| Path              | Method | Deskripsi                                                            |
	| :---------------- | :----- | :------------------------------------------------------------------- |
	| /barang           | GET    | Menampilkan data barang dengan limit 50 data terakhir                | -
	| /barang/{id}      | GET    | Menampilkan data barang per barang secara detail dengan histori stok | -
	| /barang           | POST   | Input data barang, beserta keterangan stok                           | -
	| /barang/{id}      | PUT    | Update data barang                                                   |
	| /barang/stok/{id} | PUT    | Update data stok barang saja                                         |
	| /barang/{id}      | DELETE | Hapus barang, beserta histori stok                                   |
*/
func CreateBarang(c *fiber.Ctx) error {
	type AddBarangReq struct {
		Kode       string  `json:"kode_barang"`
		Nama       string  `json:"nama_barang"`
		HargaPokok float64 `json:"harga_pokok"`
		HargaJual  float64 `json:"harga_jual"`
		Tipe       string  `json:"tipe_barang"`
		Stok       uint    `json:"stok"`
		CreateBy   string  `json:"created_by"`
		Histori    struct {
			Amount     int    `json:"amount"`
			Status     string `json:"status"`
			Keterangan string `json:"keterangan"`
		} `json:"histori_stok"`
	}
	req := new(AddBarangReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Invalid Body",
			})

	}

	barang, errCreateBarang := utils.CreateBarang(model.Barang{
		KodeBarang: req.Kode,
		Nama:       req.Nama,
		HargaPokok: req.HargaPokok,
		HargaJual:  req.HargaJual,
		TipeBarang: req.Tipe,
		Stok:       req.Stok,
		CreatedBy:  req.CreateBy,
	})

	utils.CreateHistoriBarang(&model.Details{
		ID:         barang.ID,
		KodeBarang: req.Kode,
		Nama:       req.Nama,
		HargaPokok: req.HargaPokok,
		HargaJual:  req.HargaJual,
		TipeBarang: req.Tipe,
		Stok:       req.Stok,
		CreatedBy:  req.CreateBy,
		Histori:    []model.HistoriASKM{},
	}, req.Histori.Keterangan, int(req.Stok), req.Histori.Status)

	if errCreateBarang != nil {
		logrus.Printf("Terjadi error : %s\n", errCreateBarang.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(map[string]any{
			"id":          barang.ID,
			"kode_barang": barang.KodeBarang,
		})
}

func GetBarang(c *fiber.Ctx) error {
	dataBarang, err := utils.GetBarang()
	if err != nil {
		logrus.Error("Gagal dalam mengambil list Barang: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data": dataBarang,
		},
	)
}

func GetBarangByID(c *fiber.Ctx) error {
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid ID",
			},
		)
	}

	dataBarang, err := utils.GetBarangByID(uint64(barangID))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data": dataBarang,
		},
	)
}

func UpdateBarang(c *fiber.Ctx) error {
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var updatedBarang model.Barang
	if err := c.BodyParser(&updatedBarang); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	dataBarang, err := utils.UpdateBarang(uint(barangID), updatedBarang)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update item",
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"id": dataBarang.ID,
		},
	)
}

func DeleteBarang(c *fiber.Ctx) error {
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid ID",
			},
		)
	}

	err = utils.DeleteBarang(uint64(barangID))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "Deleted Successfuly",
	},
	)
}

func UpdateStok(c *fiber.Ctx) error {
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var updatedBarang model.Barang
	if err := c.BodyParser(&updatedBarang); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	dataBarang, err := utils.UpdateBarang(uint(barangID), updatedBarang)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update item",
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"id":          dataBarang.ID,
			"kode_barang": dataBarang.KodeBarang,
		},
	)
}

//DROP TABLE `penjualans`, `models`, `historis`, `diskons`, `barangs`
