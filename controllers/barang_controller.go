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
		Nama       string  `json:"nama"`
		HargaPokok float64 `json:"harga_pokok"`
		HargaJual  float64 `json:"harga_jual"`
		Tipe       string  `json:"tipe_barang"`
		Stok       uint    `json:"stok"`
	}
	req := new(AddBarangReq)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Invalid Body",
			})

	}

	barang, errCreateBarang := utils.CreateBarang(model.Barang{
		Nama:       req.Nama,
		HargaPokok: req.HargaPokok,
		HargaJual:  req.HargaJual,
		TipeBarang: req.Tipe,
		Stok:       req.Stok,
	})

	if errCreateBarang != nil {
		logrus.Printf("Terjadi error : %s\n", errCreateBarang.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(map[string]any{
			"message": "Berhasil Menambahkan Barang",
			"Barang":  barang,
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
			"data":    dataBarang,
			"message": "Success",
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

	dataBarang, err := utils.GetBarangByID(uint(barangID))
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
			"data":    dataBarang,
			"message": "Success",
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
			"data":    dataBarang,
			"message": "Success",
		},
	)
}
