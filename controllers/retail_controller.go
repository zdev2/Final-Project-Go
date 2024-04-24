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
   | /barang           | GET    | Menampilkan data barang dengan limit 50 data terakhir                |
   | /barang/{id}      | GET    | Menampilkan data barang per barang secara detail dengan histori stok |
   | /barang           | POST   | Input data barang, beserta keterangan stok                           |
   | /barang/{id}      | PUT    | Update data barang                                                   |
   | /barang/stok/{id} | PUT    | Update data stok barang saja                                         |
   | /barang/{id}      | DELETE | Hapus barang, beserta histori stok                                   |

*/

func RouteRetail(app *fiber.App) {
	retailGroup := app.Group("/toko")

	//Barang
	retailGroup.Get("/barang", GetBarang)
	retailGroup.Get("/barang/:id", GetBarangByID)
	retailGroup.Post("/barang", CreateBarang)
	// retailGroup.Put("/barang/:id", UpdateBarang)
	// retailGroup.Put("/barang/stok/:id")
	// retailGroup.Delete("/barang/:id",)

	//Penjualan

	//Stok
}

func CreateBarang(c *fiber.Ctx) error {
	type AddBarangReq struct {
		Kode       string `json:"kode_barang"`
		Nama       string `json:"nama_barang"`
		HargaPokok int    `json:"harga_pokok"`
		HargaJual  int    `json:"harga_jual"`
		Tipe       string `json:"tipe_barang"`
		Stok       int    `json:"stok"`
	}
	req := new(AddBarangReq)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Body Not Valid",
			})

	}

	barang, errCreateBarang := utils.CreateBarang(model.Barang{})

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
				"message": "ID Not Valid",
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
