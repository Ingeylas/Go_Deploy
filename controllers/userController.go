package controllers

import (
	"TugasMID2/configs"
	"TugasMID2/middlewares"
	"TugasMID2/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Register success",
		"data":    user,
	})
}

func LoginUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	var existingUser models.User
	if err := configs.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&existingUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Email or password is wrong",
		})
	}

	tokenUser, err := middlewares.CreateToken(existingUser.ID, existingUser.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var userLoginResponse models.UserLoginResponse
	userLoginResponse.ID = existingUser.ID
	userLoginResponse.Email = existingUser.Email
	userLoginResponse.Token = tokenUser

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success",
		"data":    userLoginResponse,
	})

}
