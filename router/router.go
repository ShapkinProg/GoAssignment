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
				c.JSON(http.StatusBadRequest, models.Error("Ошибка получения сотрудников"))
				return
			}
			c.JSON(http.StatusOK, models.Success(emps))
		})

		api.POST("/employees", func(c *gin.Context) {
			var emp models.Employee
			if err := c.ShouldBindJSON(&emp); err != nil {
				log.WithError(err).Error("Ошибка парсинга JSON сотрудника")
				c.JSON(http.StatusBadRequest, models.Error("Ошибка получения данных сотрудника"))
				return
			}
			if err := service.CreateEmployee(&emp); err != nil {
				log.WithError(err).Error("Ошибка сохранения сотрудника")
				c.JSON(http.StatusBadRequest, models.Error("Ошибка сохранения сотрудника"))
				return
			}
			c.JSON(http.StatusCreated, models.Success(emp))
		})

		api.PUT("/employees/:id", func(c *gin.Context) {
			var input models.Employee
			if err := c.ShouldBindJSON(&input); err != nil {
				log.WithError(err).Error("Ошибка парсинга JSON при обновлении сотрудника")
				c.JSON(http.StatusBadRequest, models.Error("Неверный формат JSON"))
				return
			}

			updated, err := service.UpdateEmployee(c.Param("id"), &input)
			if err != nil {
				log.WithError(err).Error("Ошибка обновления сотрудника")
				c.JSON(http.StatusBadRequest, models.Error(err.Error()))
				return
			}

			c.JSON(http.StatusOK, models.Success(updated))
		})

		api.DELETE("/employees/:id", func(c *gin.Context) {
			if err := service.DeleteEmployee(c.Param("id")); err != nil {
				log.WithError(err).Error("Ошибка удаления сотрудника")
				c.JSON(http.StatusBadRequest, models.Error("Ошибка удаления сотрудника"))
				return
			}
			c.JSON(http.StatusOK, models.Success(nil))
		})

		api.GET("/departments", func(c *gin.Context) {
			depts, err := service.GetAllDepartments()
			if err != nil {
				log.WithError(err).Error("Ошибка получения отделов")
				c.JSON(http.StatusBadRequest, models.Error("Ошибка получения отделов"))
				return
			}
			c.JSON(http.StatusOK, models.Success(depts))
		})

		api.POST("/departments", func(c *gin.Context) {
			var dept models.Department
			if err := c.ShouldBindJSON(&dept); err != nil {
				log.WithError(err).Error("Ошибка парсинга JSON отдела")
				c.JSON(http.StatusBadRequest, models.Error("Ошибка чтения JSON отдела"))
				return
			}
			if err := service.CreateDepartment(&dept); err != nil {
				log.WithError(err).Error("Ошибка сохранения отдела")
				c.JSON(http.StatusInternalServerError, models.Error("Ошибка сохранения отдела"))
				return
			}
			c.JSON(http.StatusCreated, models.Success(dept))
		})

		api.PUT("/departments/:id", func(c *gin.Context) {
			var dept models.Department
			if err := c.ShouldBindJSON(&dept); err != nil {
				log.WithError(err).Error("Ошибка парсинга JSON при обновлении отдела")
				c.JSON(http.StatusBadRequest, models.Error("Ошибка JSON отдела при обновлении"))
				return
			}
			if err := service.UpdateDepartment(c.Param("id"), &dept); err != nil {
				log.WithError(err).Error("Ошибка обновления отдела")
				c.JSON(http.StatusInternalServerError, models.Error("Ошибка обновления отдела"))
				return
			}
			c.JSON(http.StatusOK, models.Success(dept))
		})

		api.DELETE("/departments/:id", func(c *gin.Context) {
			if err := service.DeleteDepartment(c.Param("id")); err != nil {
				log.WithError(err).Error("Ошибка удаления отдела")
				c.JSON(http.StatusInternalServerError, models.Error("Ошибка удаления отдела"))
				return
			}
			c.JSON(http.StatusOK, models.Success(nil))
		})
	}
}
