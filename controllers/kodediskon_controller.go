package controllers

import (
	"kelompok3/toko-retail/model"
	"kelompok3/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func CreateKodeDiskon(c *fiber.Ctx) error {
	type AddKodeDiskon struct {
		Kode_diskon string  `json:"kode_diskon"`
		Amount      float64 `json:"amount"`
		Type        string  `json:"type"`
	}

	req := new(AddKodeDiskon)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Invalid Body",
			})

	}

	diskon, errDiskon := utils.CreateKodeDiskon(model.Diskon{})

	if errDiskon != nil {
		logrus.Printf("Terjadi error : %s\n", errDiskon.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(map[string]any{
			"message": "Berhasil Menambahkan Kode Diskon",
			"data":    diskon,
		})

}

func GetKodeDiskon(c *fiber.Ctx) error {
	dataDiskon, err := utils.GetDiskon()

	if err != nil {
		logrus.Error("Gagal dalam mengambil list Diskon: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    dataDiskon,
			"message": "Success",
		},
	)
}

func GetByCode(c *fiber.Ctx) error {
	DiskonCode := c.Params("kode-diskon")

	dataDiskon, err := utils.GetDiskonByCode(string(DiskonCode))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "Discount Code not found",
				},
			)
		}
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    dataDiskon,
			"message": "Success",
		},
	)
}

func GetDiskonByID(c *fiber.Ctx) error {
	DiskonID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid ID",
			},
		)
	}

	dataDiskon, err := utils.GetDiskonByID(uint(DiskonID))
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
			"data":    dataDiskon,
			"message": "Success",
		},
	)
}

func UpdateCode(c *fiber.Ctx) error {
	DiskonID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var updatedDiskon model.Diskon
	if err := c.BodyParser(&updatedDiskon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	dataDiskon, err := utils.UpdateDiskon(uint(DiskonID), updatedDiskon)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update item",
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    dataDiskon,
			"message": "Success",
		},
	)
}
