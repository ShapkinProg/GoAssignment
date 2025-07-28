package service

import (
	"GoAssignment/db"
	"GoAssignment/models"
	"errors"
	"fmt"
	"strings"
)

func GetAllEmployees() ([]models.Employee, error) {
	var emps []models.Employee
	// err := db.DB.Preload("Department").Find(&emps).Error
	err := db.DB.Preload("Department").Where("is_deleted = ?", false).Find(&emps).Error
	return emps, err
}

func CreateEmployee(emp *models.Employee) error {
	emp.Name = strings.TrimSpace(emp.Name)
	if err := Validate(emp); err != nil {
		return err
	}
	return db.DB.Create(emp).Error
}

func UpdateEmployee(id string, input *models.Employee) (*models.Employee, error) {
	var emp models.Employee
	if err := db.DB.First(&emp, id).Error; err != nil {
		return nil, err
	}

	updated := false

	if input.Name != "" {
		emp.Name = input.Name
		updated = true
	}
	if input.Position != "" {
		emp.Position = input.Position
		updated = true
	}
	if input.DepartmentID != 0 && input.DepartmentID != emp.DepartmentID {
		emp.DepartmentID = uint(input.DepartmentID)
		updated = true
	}

	if !updated {
		return nil, errors.New("нет данных для обновления")
	}

	if err := Validate(&emp); err != nil {
		return nil, err
	}

	if err := db.DB.Save(&emp).Error; err != nil {
		return nil, err
	}

	var updatedEmp models.Employee
	if err := db.DB.Preload("Department").First(&updatedEmp, emp.ID).Error; err != nil {
		return nil, err
	}

	return &updatedEmp, nil
}

func DeleteEmployee(id string) error {
	// return db.DB.Delete(&models.Employee{}, id).Error
	return db.DB.Model(&models.Employee{}).Where("id = ?", id).Update("is_deleted", true).Error
}

func GetAllDepartments() ([]models.Department, error) {
	var depts []models.Department
	err := db.DB.Where("is_deleted = ?", false).Find(&depts).Error
	return depts, err
}

func CreateDepartment(dept *models.Department) error {
	dept.Name = strings.TrimSpace(dept.Name)
	if dept.Name == "" {
		return errors.New("название отдела не может быть пустым")
	}
	return db.DB.Create(dept).Error
}

func UpdateDepartment(id string, input *models.Department) error {
	var dept models.Department
	if err := db.DB.First(&dept, id).Error; err != nil {
		return err
	}
	if strings.TrimSpace(input.Name) == "" {
		return errors.New("название отдела не может быть пустым")
	}
	dept.Name = strings.TrimSpace(input.Name)
	return db.DB.Save(&dept).Error
}

func DeleteDepartment(id string) error {
	// return db.DB.Delete(&models.Department{}, id).Error
	return db.DB.Model(&models.Department{}).Where("id = ?", id).Update("is_deleted", true).Error
}

func Validate(e *models.Employee) error {
	if e.Name != "" {
		parts := len(strings.Fields(e.Name))
		if parts != 3 {
			return errors.New("ФИО должно содержать 3 слова")
		}
	}
	if e.Position != "" && len(e.Position) < 3 {
		return errors.New("должность слишком короткая, минимум 3 символа")
	}
	var count int64
	err := db.DB.Model(&models.Department{}).Where("id = ? AND is_deleted = false", e.DepartmentID).Count(&count).Error
	if err != nil {
		return errors.New("ошибка при проверке отдела")
	}
	if count == 0 {
		return fmt.Errorf("отдел с ID %d и именем %s не найден", e.DepartmentID, e.Department.Name)
	}
	return nil
}
