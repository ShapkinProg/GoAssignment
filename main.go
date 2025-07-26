package main

import (
	"GoAssignment/db"
	"GoAssignment/router"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	db.Init()
	r := gin.Default()
	r.Static("/assets", "./frontend/dist/assets")
	router.SetupRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	log.Info("Сервер запущен на порту 8080")
	err := r.Run(":8080")
	if err != nil {
		log.WithError(err).Fatal("Ошибка запуска сервера")
	}
}
