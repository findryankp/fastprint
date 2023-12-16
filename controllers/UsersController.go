package controllers

import (
	"fastprint/configs"
	"fastprint/helpers"
	"fastprint/models"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var user models.Users
	if err := c.Bind(&user); err != nil {
		return c.JSON(400,helpers.ResponseFail(err.Error()))
	}

	query := configs.DB.Create(&user)
	if query.Error != nil {
		return c.JSON(404,helpers.ResponseFail("not found"))
	}

	return c.JSON(201,helpers.ResponseSuccess("Registered succesfully",user))
}

func Login(c echo.Context) error {
	var user models.Users
	var request models.Users

	if err := c.Bind(&request); err != nil {
		return c.JSON(400,helpers.ResponseFail(err.Error()))
	}

	record := configs.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		return c.JSON(400,helpers.ResponseFail(record.Error.Error()))
	}

	credentialError := helpers.CheckPassword(request.Password, user.Password)
	if !credentialError {
		return c.JSON(401,helpers.ResponseFail("Invalid Credential"))
	}

	tokenString, err := helpers.GenerateToken(int(user.ID))
	if err != nil {
		return c.JSON(400,helpers.ResponseFail(err.Error()))
	}

	response := map[string]string{
		"token": tokenString,
	}

	return c.JSON(200,helpers.ResponseSuccess("-",response))
}
