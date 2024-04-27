package controllers

import (
	"kelompok3/toko-retail/model"
	"kelompok3/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func InsertPenjualanData(c *fiber.Ctx) error {
	type AddPenjualanReq struct {
		NamaPembeli string  `json:"nama"`
		ID          uint64  `json:"id"`
		Subtotal    float64 `json:"subtotal"`
		Total       float64 `json:"total"`
	}
	req := new(AddPenjualanReq)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Invalid Body",
			})

	}

	penjualan, errInsertPenjualan := utils.InsertPenjualanData(model.Penjualan{
		Nama_pembeli: req.NamaPembeli,
		ID:           req.ID,
		Subtotal:     req.Subtotal,
		Total:        req.Total,
	})

	if errInsertPenjualan != nil {
		logrus.Printf("Terjadi error : %s\n", errInsertPenjualan.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(map[string]any{
			"message": "Berhasil Menambahkan Barang",
			"Barang":  penjualan,
		})
}

func GetPenjualan(c *fiber.Ctx) error {
	dataPenjualan, err := utils.GetPenjualan()
	if err != nil {
		logrus.Error("Gagal dalam mengambil list Penjualan :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    dataPenjualan,
			"message": "Success",
		},
	)
}

func GetPenjualanByID(c *fiber.Ctx) error {
	penjualanID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid ID",
			},
		)
	}

	dataPenjualan, err := utils.GetPenjualanByID(uint64(penjualanID))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID Not Found",
				},
			)
		}
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    dataPenjualan,
			"message": "Success",
		},
	)
}
