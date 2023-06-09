package database

import (
	"Hacktiv10JWT/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "17160100juni"
	port     = 5432
	dbname   = "simple-api-jwt"

	db  *gorm.DB
	err error
)

func StartDB() {
	fmt.Println("connecting to database....")
	config := fmt.Sprintf("host= %s user= %s password= %s port= %d dbname= %s sslmode= disable",
		host, user, password, port, dbname)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	fmt.Println("connected to database")

	// Log the database object to check if it's nil
	fmt.Printf("DB object: %v\n", db)

	// Auto-migrate the models
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}

func PingDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if sqlDB != nil {
		err := sqlDB.Close()
		if err != nil {
			log.Fatalf("error closing database connection: %v", err)
		}
	}
}
