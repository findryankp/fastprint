package configs

import (
	"fmt"
	"os"
	"fastprint/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	USER := os.Getenv("DBUSERNAME")
	PASS := os.Getenv("DBPASS")
	HOST := os.Getenv("DBHOST")
	PORT := os.Getenv("DBPORT")
	DBNAME := os.Getenv("DBNAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(
		&models.Produk{},
		&models.Kategori{},
		&models.Status{},
		&models.Users{},
	)
	fmt.Println("Migration done.")
}
	
