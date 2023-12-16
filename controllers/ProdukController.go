package controllers

import (
	"fastprint/configs"
	"fastprint/helpers"
	"fastprint/models"
	"github.com/labstack/echo/v4"
)

func Produk(c echo.Context) error {
	var data []models.Produk
	if err := configs.DB.Find(&data); err.Error != nil {
		return c.JSON(500, helpers.ResponseFail(err.Error.Error()))
	}
	return c.JSON(200, helpers.ResponseSuccess("-", data))
}

func ProdukById(c echo.Context) error {
	var data models.Produk
	if err := configs.DB.First(&data, c.Param("id")); err.Error != nil {
		return c.JSON(500, helpers.ResponseFail(err.Error.Error()))
	}
	return c.JSON(200, helpers.ResponseSuccess("-", data))
}

func ProdukCreate(c echo.Context) error {
	var request models.ProdukRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(500, helpers.ResponseFail("-"))
	}

	data := models.RequestToModelProduk(request)
	if err := configs.DB.Create(&data); err.Error != nil {
		return c.JSON(500, helpers.ResponseFail(err.Error.Error()))
	}

	return c.JSON(201, helpers.ResponseSuccess("data created", nil))
}

func ProdukUpdate(c echo.Context) error {
	var request models.ProdukRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(500, helpers.ResponseFail("-"))
	}

	data := models.RequestToModelProduk(request)
	if err := configs.DB.Where("id = ?", c.Param("id")).Updates(&data); err.Error != nil {
		return c.JSON(500, helpers.ResponseFail(err.Error.Error()))
	}

	return c.JSON(200, helpers.ResponseSuccess("data updated", nil))
}

func ProdukDelete(c echo.Context) error {
	var data models.Produk
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&data); err.Error != nil {
		return c.JSON(404, helpers.ResponseFail("data not found"))
	}

	configs.DB.Delete(&data)

	return c.JSON(200, helpers.ResponseSuccess("data deleted", nil))
}
	