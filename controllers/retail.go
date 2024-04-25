package controllers

import "github.com/gofiber/fiber/v2"

func RouteRetail(app *fiber.App) {
	retailGroup := app.Group("/toko")

	//Barang
	retailGroup.Get("/barang", GetBarang)
	retailGroup.Get("/barang/:id", GetBarangByID)
	retailGroup.Post("/barang", CreateBarang)
	retailGroup.Put("/barang/:id", UpdateBarang)
	// retailGroup.Put("/barang/stok/:id")
	// retailGroup.Delete("/barang/:id",)

	//Penjualan
	// retailGroup.Get("/penjualan")
	// retailGroup.Get("/penjualan/:id")
	// retailGroup.Post("/penjualan")

	//Kode Diskon
	retailGroup.Get("/kode-diskon", GetKodeDiskon)
	retailGroup.Get("/kode-diskon/:id", GetDiskonByID)
	retailGroup.Get("/kode-diskon/get-by-code", GetByCode)
	retailGroup.Post("/kode-diskon", CreateKodeDiskon)
}
