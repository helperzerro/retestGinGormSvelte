// models/employee.go
package models

type Employee struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"size:255" json:"name"`
	NPWP    string `gorm:"size:50" json:"npwp"`
	Address string `gorm:"size:255" json:"address"`
}
