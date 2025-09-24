package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// sesuaikan dengan user/password/db kamu
	host 		:= "your_host"
	user 		:= "your_user"
	password 	:= "your_password"
	dbName 		:= "your_dbname"
	port 		:= "your_port"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database... :", err)
	}

	fmt.Println("Koneksi ke database berhasil!")
}
