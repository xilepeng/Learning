package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// create  database  test character set utf8mb4 collate utf8mb4_unicode_ci;
// use test;

// 定义模型
type User struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	dsn := "root:Yizhili80@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         nil,
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 单数表名
	})
	if err != nil {
		panic(err)
	}

	// 数据迁移：把模型与数据库表对应起来
	db.AutoMigrate(&User{})

	// 创建
	u := User{Name: "X", Age: 18}

	db.Create(&u)
	db.Debug()
}
