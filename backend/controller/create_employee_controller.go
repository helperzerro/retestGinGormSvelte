package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helperzerro/ginGormSVELTE/models"
	"gorm.io/gorm"
)

func CreateEmployeeController(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var employee models.Employee

		// Bind JSON dari request body ke struct
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
			return
		}

		// Simpan ke database
		if err := db.Create(&employee).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Balikin response sukses
		c.JSON(http.StatusCreated, gin.H{"message": "Data berhasil disimpan", "data": employee})
	}
}
