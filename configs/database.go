package configs

import (
	"TugasMID2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:12345678@tcp(localhost:3306)/tugasmid2?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

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
