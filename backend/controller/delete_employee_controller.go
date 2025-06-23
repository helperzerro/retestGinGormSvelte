package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/helperzerro/ginGormSVELTE/models"
	"gorm.io/gorm"
)

// DeleteEmployeeHandler untuk menghapus data karyawan berdasarkan ID
func DeleteEmployeeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil ID dari path param
		idStr := c.Param("id")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
			return
		}

		// Mengonversi ID menjadi integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Hapus record berdasarkan ID
		if err := db.Delete(&models.Employee{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee", "details": err.Error()})
			return
		}

		// Jika berhasil, kembalikan status 200 atau 204
		c.JSON(http.StatusOK, gin.H{"message": "Employee successfully deleted"})
	}
}
