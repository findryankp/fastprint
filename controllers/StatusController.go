package controllers

import (
	"fastprint/configs"
	"fastprint/helpers"
	"fastprint/models"
	"github.com/labstack/echo/v4"
)

func Status(c echo.Context) error {
	var data []models.Status
	if err := configs.DB.Find(&data); err.Error != nil {
		return c.JSON(500, helpers.ResponseFail(err.Error.Error()))
	}
	return c.JSON(200, helpers.ResponseSuccess("-", data))
}

func StatusById(c echo.Context) error {
	var data models.Status
	if err := configs.DB.First(&data, c.Param("id")); err.Error != nil {
		return c.JSON(500, helpers.ResponseFail(err.Error.Error()))
	}
	return c.JSON(200, helpers.ResponseSuccess("-", data))
}

func StatusCreate(c echo.Context) error {
	var request models.StatusRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(500, helpers.ResponseFail("-"))
	}

	data := models.RequestToModelStatus(request)
	if err := configs.DB.Create(&data); err.Error != nil {
		return c.JSON(500, helpers.ResponseFail(err.Error.Error()))
	}

	return c.JSON(201, helpers.ResponseSuccess("data created", nil))
}

func StatusUpdate(c echo.Context) error {
	var request models.StatusRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(500, helpers.ResponseFail("-"))
	}

	data := models.RequestToModelStatus(request)
	if err := configs.DB.Where("id = ?", c.Param("id")).Updates(&data); err.Error != nil {
		return c.JSON(500, helpers.ResponseFail(err.Error.Error()))
	}

	return c.JSON(200, helpers.ResponseSuccess("data updated", nil))
}

func StatusDelete(c echo.Context) error {
	var data models.Status
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&data); err.Error != nil {
		return c.JSON(404, helpers.ResponseFail("data not found"))
	}

	configs.DB.Delete(&data)

	return c.JSON(200, helpers.ResponseSuccess("data deleted", nil))
}
	