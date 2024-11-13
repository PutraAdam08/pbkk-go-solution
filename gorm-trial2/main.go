package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func ConnectDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect db")
	}
	return db

}

func viewAllProduct(db *gorm.DB) {
	var products []Product
	result := db.Find(&products)

	if result.Error != nil {
		fmt.Println("Error retrieving products:", result.Error)
		return
	}

	sort := 1
	for _, product := range products {
		fmt.Printf("%d. Code 	: %s\n   Price	: %d\n", sort, product.Code, product.Price)
		sort++
	}
}

func addProduct(db *gorm.DB) {
	var Code string
	var Price uint
	fmt.Printf("===Insert Product===\n")
	fmt.Printf("Code	: ")
	fmt.Scan(Code)
	fmt.Printf("\nPrice	: ")
	fmt.Scan(Price)
	db.Create(&Product{Code: Code, Price: Price})
}

func editProduct(db *gorm.DB) {

}

func deleteProduct(db *gorm.DB) {
	var code string
	fmt.Printf("Input product code: ")
	fmt.Scan(&code)

	result := db.Where("code = ?", code).Delete(&Product{})

	if result.Error != nil {
		fmt.Println("Error deleting product:", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		fmt.Println("No product found with the specified Code.")
	} else {
		fmt.Println("Product deleted successfully.")
	}
}

func main() {
	db := ConnectDB()
	var option int

	fmt.Println("=== Welcome to Product Management ===")
	fmt.Printf("What you want to do? :\n1. View product\n2.Add Product\n3. Edit Product\n4. Delete Product\n")

	fmt.Printf("Command: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		viewAllProduct(db)
	case 2:
		addProduct(db)
	case 3:
		editProduct(db)
	case 4:
		deleteProduct(db)
	}
}
