package controllers

import "kelompok3/toko-retail/model"

func RouteRetail(app *fiber.App) {
  retailGroup := app.group("/toko")
  retailGroup.Get("/barang", GetBarang)
  retailGroup.Get("/barang/:id", GetBarangByID)
  retailGroup.Post("/barang", CreateBarang)
  retailGroup.Put("/barang/:id", UpdateBarang)
  retailGroup.Put("/barang/stok/:id")
  retailGroup.Delete("/barang/:id",
}

func CreateBarang(app *fiber.app) error {
  type AddBarangReq struct {
    Kode string `json:"kode_barang"`
    Nama string `json:"nama_barang"`
    HargaPokok int `json:"harga_pokok"`
    HargaJual int `json:"harga_jual"`
    Tipe string `json:"tipe_barang"`
    Stok int `json:"stok"`
    
  }
  req := new(AddBarangReq)

  if err := c.BodyParser(req); err != nil {
    return c.Status(fiber.StatusBadRequest).
    JSON(map[string]any{
      "message": "Body Not Valid",
    })
    
    isValid && err != nil {
      return c.Status(fiber.StatusBadRequest).
      JSON(map[string]any{
        "message": err.Error(),
      })
      errCreateBarang := utils.CreateBarang(model.Barang{})

      if errCreateBarang != nil {
        logrus.Printf("Terjadi error : &s\n", errCreateBarang.Error())
        return c.Status(fiber.StatusInternalServerError).
        JSON(map[string]any{
          "message"; "Server Error",
        })

        return c.Status(fiber.StatusOK).
        JSON(map[string]any{
          "message": "Berhasil Menambahkan Barang",
          "Barang": model.Barang
        })
      }
    } 
  }
}