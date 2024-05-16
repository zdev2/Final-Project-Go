package controllers

import "github.com/gofiber/fiber/v2"

func RouteRetail(app *fiber.App) {
	retailGroup := app.Group("/toko")

	//Barang
	retailGroup.Get("/barang", GetBarang)
	retailGroup.Get("/barang/:id", GetBarangByID)
	retailGroup.Post("/barang", CreateBarang)
	retailGroup.Put("/barang/:id", UpdateBarang)
	// retailGroup.Put("/barang/stok/:id", )
	retailGroup.Delete("/barang/:id", DeleteBarang)

	//Penjualan
	retailGroup.Get("/penjualan", GetPenjualan)
	retailGroup.Get("/penjualan/:id")
	retailGroup.Post("/penjualan", InsertPenjualanData)

	//Kode Diskon
	retailGroup.Get("/kode-diskon", GetKodeDiskon)
	retailGroup.Get("/kode-diskon/get-by-code", GetByCode)
	retailGroup.Get("/kode-diskon/:id", GetDiskonByID)
	retailGroup.Post("/kode-diskon", CreateKodeDiskon)
	retailGroup.Put("/kode-diskon/:id", UpdateCode)
	retailGroup.Delete("/kode-diskon/:id", DeleteKode)
}
