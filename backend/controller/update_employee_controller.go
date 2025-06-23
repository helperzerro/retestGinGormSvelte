package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/helperzerro/ginGormSVELTE/models"
	"gorm.io/gorm"
)

// Get Data By Id
func GetEmployeeByIDController(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan ID dari URL parameter
		idStr := c.Param("id")

		// Mengonversi ID menjadi integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
			return
		}

		// Mencari karyawan berdasarkan ID
		var employee models.Employee
		if err := db.First(&employee, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan tidak ditemukan"})
			return
		}

		// Jika karyawan ditemukan, kirimkan data karyawan
		c.JSON(http.StatusOK, employee)
	}
}

// Controller untuk memperbarui data karyawan berdasarkan ID
func UpdateEmployeeController(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
			return
		}

		var input models.Employee
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
			return
		}

		if err := db.Model(&models.Employee{}).Where("id = ?", id).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data karyawan berhasil diupdate"})
	}
}

