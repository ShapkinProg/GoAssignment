package db

import (
	"GoAssignment/models"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	local := flag.Bool("local", false, "Локальное подключение")
	flag.BoolVar(local, "l", false, "Синоним --local")
	flag.Parse()
	godotenv.Load()
	var dsn string
	if *local {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST_LOCAL"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	} else {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST_DOCKER"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	}

	var err error
	for {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.WithError(err).Error("Ошибка подключения к БД, пробуем снова через 5 секунд")
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	log.Info("Успешное подключение к БД")

	err = DB.AutoMigrate(&models.Department{}, &models.Employee{})
	if err != nil {
		log.WithError(err).Fatal("Ошибка миграции БД")
	}
	log.Info("Миграция завершена")
	// DB.Exec(`ALTER TABLE employees DROP COLUMN IF EXISTS deleted_at`)
	// DB.Exec(`ALTER TABLE departments DROP COLUMN IF EXISTS deleted_at`)
	seedData()
}

func seedData() {
	var count int64
	DB.Model(&models.Department{}).Count(&count)
	if count > 0 {
		log.Info("Сидинг пропущен")
		return
	}

	depts := []models.Department{
		{Name: "Отдел разработки"},
		{Name: "Отдел продаж"},
		{Name: "HR"},
	}
	DB.Create(&depts)

	DB.Where("name = ?", "Отдел разработки").First(&depts[0])
	DB.Where("name = ?", "Отдел продаж").First(&depts[1])
	DB.Where("name = ?", "HR").First(&depts[2])

	employees := []models.Employee{
		{Name: "Иван Иванов Иванович", Position: "Backend разработчик", DepartmentID: depts[0].ID},
		{Name: "Пётр Петров Сергеевич", Position: "Frontend разработчик", DepartmentID: depts[0].ID},
		{Name: "Анна Смирнова Олеговна", Position: "Менеджер по продажам", DepartmentID: depts[1].ID},
		{Name: "Мария Козлова Дмитриевна", Position: "HR-специалист", DepartmentID: depts[2].ID},
	}
	DB.Create(&employees)
	log.Info("Сидинг завершён")
}
