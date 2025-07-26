package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Employees []Employee
}

type Employee struct {
	gorm.Model
	Name         string
	DepartmentID uint
	Position     string
	Department   Department
}
