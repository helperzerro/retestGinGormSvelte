package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helperzerro/ginGormSVELTE/models"
	"gorm.io/gorm"
)

func IndexEmployeeController(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var employees []models.Employee

		// Ambil data dari database
		if err := db.Find(&employees).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Kirim data sebagai JSON
		c.JSON(http.StatusOK, gin.H{
			"employees": employees,
		})
	}
}
