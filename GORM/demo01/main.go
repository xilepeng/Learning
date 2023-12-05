package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint `gorm:"column:id;type:uint;primaryKey;autoIncrement"`
	Code  string
	Price uint
}

// create  database  test character set utf8mb4 collate utf8mb4_unicode_ci;
// use test;

func main() {
	dsn := "root:Yizhili80@tcp(127.0.0.1:3306)/test"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})
	insertProduct := &Product{Code: "D42", Price: 100}
	db.Create(insertProduct)

	fmt.Printf("insert ID: %d, Code: %s, Price: %d\n",
		insertProduct.ID, insertProduct.Code, insertProduct.Price)

	readProduct := &Product{}
	db.First(&readProduct, "code = ?", "D42") // find product with code D42

	fmt.Printf("read ID: %d, Code: %s, Price: %d\n",
		readProduct.ID, readProduct.Code, readProduct.Price)
}

// ➜  demo01 git:(main) ✗ go run main.go
// insert ID: 1, Code: D42, Price: 100
// read ID: 1, Code: D42, Price: 100
