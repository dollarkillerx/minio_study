package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

type Test struct {
	gorm.Model
	File []byte `gorm:"type:LONGBLOB"`
}



func getDB() *gorm.DB {
	dsn := "root:@tcp(192.168.88.14:4000)/big?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&Test{}); err != nil {
		log.Fatalln(err)
	}

	return db
}



func main() {
	db := getDB()

	file, err := ioutil.ReadFile("rust_book.pdf")
	if err != nil {
		log.Fatalln(err)
	}
	f1 := Test{File: file}
	if err := db.Create(&f1).Error; err != nil {
		log.Fatalln(err)
	}
	log.Println("SUCCESS")
}
