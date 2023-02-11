package database

import (
	"log"
	"os"

	"github.com/firdisml/go-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Instance struct {
	Db *gorm.DB
}

var Database Instance

func Connect() {
	dsn := "host=containers-us-west-189.railway.app user=postgres password=1aOUGSpULEROiwM8foxJ dbname=railway port=7713 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = Instance{Db: db}
}
