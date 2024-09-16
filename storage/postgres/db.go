package postgres

import (
	"log"

	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnToDb() (*gorm.DB, error) {

	dsn := "host=0.0.0.0 user=jasur password=1001 dbname=kafka_todo_app port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		log.Println("error on ConnToDB(). Error:", err)
		return nil, err
	}

	return db, err
}

func AoutoMigrate(db *gorm.DB) error {

	// if err := db.AutoMigrate(
	// 	&models.User{},
	// ); err != nil {
	// 	log.Println("error on AoutoMigrate(). Error:", err)
	// 	return err
	// }
	if err := db.AutoMigrate(
		&models.Todo{},
	); err != nil {
		log.Println("error on AoutoMigrate(). Error:", err)
		return err
	}

	return nil
}
