package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:909923@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=true&loc=Local"))
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// db.AutoMigrate(&Product{})
	//db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Message{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})
	//user := &models.UserBasic{}
	//user.Name = "小强"
	//db.Create(user)
	//
	//// Read
	//// var product Product
	//// fmt.Println(db.First(&product, 1))    // find product with integer primary key
	//// db.First(&product, "code = ?", "D42") // find product with code D42
	//fmt.Println(db.First(user, 1))
	//
	//// Update - update product's price to 200
	//// db.Model(&product).Update("Price", 200)
	//db.Model(user).Update("PassWord", "1234")
	// Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(&product, 1)
}
