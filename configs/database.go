package configs

import (
	"TugasMID2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"fmt"
)

var DB *gorm.DB

func InitDB() {
	// dsn := "root:12345678@tcp(localhost:3306)/tugasmid2?charset=utf8mb4&parseTime=True&loc=Local"
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")


	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	InitMigrate()
}

func InitMigrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}

	if err := DB.AutoMigrate(&models.Package{}); err != nil {
		panic(err.Error())
	}
}
