package main

import (
	//"fmt"

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

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D11", Price: 25})

	var product Product
	//db.First(&product, "code = ?", "D11")
	//db.Model(&product).Update("Price", 200)

	// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//db.Delete(&product, 1)
	db.First(&product, 1)
	//fmt.Printf("Code: %s\nPrice: %d", product.Code, product.Price)

}
