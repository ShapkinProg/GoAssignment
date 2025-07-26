package router

import (
	"GoAssignment/models"
	"GoAssignment/service"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/employees", func(c *gin.Context) {
			emps, err := service.GetAllEmployees()
			if err != nil {
				log.WithError(err).Error("Ошибка получения сотрудников")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка"})
				return
			}
			c.JSON(http.StatusOK, emps)
		})

		api.POST("/employees", func(c *gin.Context) {
			var emp models.Employee
			if err := c.ShouldBindJSON(&emp); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
				return
			}
			if err := service.CreateEmployee(&emp); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать"})
				return
			}
			c.JSON(http.StatusCreated, emp)
		})

		api.PUT("/employees/:id", func(c *gin.Context) {
			id := c.Param("id")
			var emp models.Employee
			if err := c.ShouldBindJSON(&emp); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
				return
			}
			if err := service.UpdateEmployee(id, &emp); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления"})
				return
			}
			c.JSON(http.StatusOK, emp)
		})

		api.DELETE("/employees/:id", func(c *gin.Context) {
			if err := service.DeleteEmployee(c.Param("id")); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления"})
				return
			}
			c.Status(http.StatusNoContent)
		})

		api.GET("/departments", func(c *gin.Context) {
			depts, err := service.GetAllDepartments()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка"})
				return
			}
			c.JSON(http.StatusOK, depts)
		})

		api.POST("/departments", func(c *gin.Context) {
			var dept models.Department
			if err := c.ShouldBindJSON(&dept); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
				return
			}
			if err := service.CreateDepartment(&dept); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания"})
				return
			}
			c.JSON(http.StatusCreated, dept)
		})

		api.PUT("/departments/:id", func(c *gin.Context) {
			var dept models.Department
			if err := c.ShouldBindJSON(&dept); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
				return
			}
			if err := service.UpdateDepartment(c.Param("id"), &dept); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления"})
				return
			}
			c.JSON(http.StatusOK, dept)
		})

		api.DELETE("/departments/:id", func(c *gin.Context) {
			if err := service.DeleteDepartment(c.Param("id")); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления"})
				return
			}
			c.Status(http.StatusNoContent)
		})
	}
}
