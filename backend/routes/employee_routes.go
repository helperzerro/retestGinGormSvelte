package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helperzerro/ginGormSVELTE/controller"
	"github.com/helperzerro/ginGormSVELTE/middleware" // Import middleware
	"gorm.io/gorm"
)

func EmployeeRoutes(router *gin.Engine, db *gorm.DB) {
    router.GET("/", controller.EmployeeController())

    // Rute yang tidak memerlukan autentikasi
    router.POST("/auth/register", controller.Register(db))
    router.POST("/auth/login", controller.Login(db))

    // Rute yang memerlukan autentikasi
    authorized := router.Group("/")
    authorized.Use(middleware.AuthMiddleware()) // Gunakan middleware di grup ini

    // Get All Data
    authorized.GET("/index_employee", controller.IndexEmployeeController(db))

    // Create Data
    authorized.POST("/create_employee", controller.CreateEmployeeController(db))

    // Update Data
    authorized.GET("/update_employee/:id", controller.GetEmployeeByIDController(db))
    authorized.PATCH("/update_employee/:id", controller.UpdateEmployeeController(db))

    // Delete Data
    authorized.DELETE("/delete_employee/:id", controller.DeleteEmployeeHandler(db))

    // Rute untuk menyajikan halaman HTML yang memerlukan autentikasi
    authorized.GET("/dashboard", func(c *gin.Context) {
        c.HTML(http.StatusOK, "dashboard.html", nil) // Ganti dengan nama file HTML Anda
    })
}