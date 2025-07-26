package service

import (
	"GoAssignment/db"
	"GoAssignment/models"
)

func GetAllEmployees() ([]models.Employee, error) {
	var emps []models.Employee
	err := db.DB.Preload("Department").Find(&emps).Error
	return emps, err
}

func CreateEmployee(emp *models.Employee) error {
	return db.DB.Create(emp).Error
}

func UpdateEmployee(id string, input *models.Employee) error {
	var emp models.Employee
	if err := db.DB.First(&emp, id).Error; err != nil {
		return err
	}
	emp.Name = input.Name
	emp.Position = input.Position
	emp.DepartmentID = input.DepartmentID
	return db.DB.Save(&emp).Error
}

func DeleteEmployee(id string) error {
	return db.DB.Delete(&models.Employee{}, id).Error
}

func GetAllDepartments() ([]models.Department, error) {
	var depts []models.Department
	err := db.DB.Find(&depts).Error
	return depts, err
}

func CreateDepartment(dept *models.Department) error {
	return db.DB.Create(dept).Error
}

func UpdateDepartment(id string, input *models.Department) error {
	var dept models.Department
	if err := db.DB.First(&dept, id).Error; err != nil {
		return err
	}
	dept.Name = input.Name
	return db.DB.Save(&dept).Error
}

func DeleteDepartment(id string) error {
	return db.DB.Delete(&models.Department{}, id).Error
}
