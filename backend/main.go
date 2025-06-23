package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	database "github.com/helperzerro/ginGormSVELTE/config"
	"github.com/helperzerro/ginGormSVELTE/routes"
)

func main() {
	db:=database.ConnectDB()

	r := gin.Default()
	
	// Aktifkan CORS
	r.Use(cors.Default())
	
	// Tambahkan logging sebelum server dijalankan
	log.Println("🚀 Server running at http://localhost:8080")
	
	routes.EmployeeRoutes(r,db)

	r.Run(":8080") // Jalanin di localhost:8080
}
