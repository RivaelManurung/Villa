package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connection() {
	const con = "root@tcp(localhost)/villa?charset=utf8mb4&parseTime=True&loc=Local"
	abc := con
	db, err := gorm.Open(mysql.Open(abc), &gorm.Config{})

	if err != nil {
		panic(" gagal koneksi")
	}
	DB = db
	fmt.Println("Berhasil konek ke database")
}