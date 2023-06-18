package connection

import (
	"boilerplate-echo-gorm-postgres/db/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDB() *gorm.DB {
	dsn := "host=postgres user=leyla password=leyla dbname=todo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//	db.LogMode(true)
	if err != nil {
		fmt.Println("Error Connection Database")
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Connected to Database ...")

	db.AutoMigrate(&models.Task{})

	return db

}
