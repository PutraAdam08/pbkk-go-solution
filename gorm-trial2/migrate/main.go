package main

import (
	//"fmt"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect db")
	}

	products := []*Product{
		{Code: "C21", Price: 10},
		{Code: "G71", Price: 35},
		{Code: "A12", Price: 40},
		{Code: "B69", Price: 11},
		{Code: "C42", Price: 13},
		{Code: "G30", Price: 15},
	}

	result := db.Create(products)

	if result.Error != nil {
		fmt.Println("Error inserting product:", result.Error)
	}

}
