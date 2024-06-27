package controllers

import (
	"TugasMID2/configs"
	"TugasMID2/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllPackages(c echo.Context) error {
	if err := configs.DB.Find(&models.Package{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get all packages",
		"data":    models.Package{},
	})
}

func GetPackageByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var packageData models.Package

	if err := configs.DB.Where("id = ?", id).First(&packageData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get package by ID",
		"data":    packageData,
	})

}

func CreatePackage(c echo.Context) error {
	var packageData models.Package
	if err := c.Bind(&packageData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := configs.DB.Create(&packageData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Create package success",
		"data":    packageData,
	})
}

func UpdatePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var packageData models.Package

	if err := configs.DB.Where("id = ?", id).First(&packageData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := c.Bind(&packageData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := configs.DB.Save(&packageData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Update package success",
		"data":    packageData,
	})
}

func DeletePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var packageData models.Package

	if err := configs.DB.Where("id = ?", id).First(&packageData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := configs.DB.Delete(&packageData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Delete package success",
		"data":    packageData,
	})
}
