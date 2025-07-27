// package models

// import "gorm.io/gorm"

// type Department struct {
// 	gorm.Model
// 	Name      string `gorm:"not null"`
// 	Employees []Employee
// }

//	type Employee struct {
//		gorm.Model
//		Name         string
//		DepartmentID uint
//		Position     string
//		Department   Department
//	}
package models

import (
	"time"
)

type Department struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool `gorm:"default:false"`
	Employees []Employee
}

type Employee struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Position     string `gorm:"not null"`
	DepartmentID uint   `gorm:"not null"`
	Department   Department
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsDeleted    bool `gorm:"default:false"`
}
