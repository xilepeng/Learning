package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

// create database test charset = utf8mb4;
// create  database  test character set utf8mb4 collate utf8mb4_unicode_ci;
// use test;

func main() {
	dsn := "root:Yizhili80@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&UserInfo{})

	// 创建
	u1 := UserInfo{ID: 1, Name: "X", Gender: "男", Hobby: "Coding"}
	db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("查询 u:%#v\n", u)

	// 更新
	db.Model(&u).Update("hobby", "algo")
	fmt.Printf("更新 u:%#v\n", u)

	// db.Model(&u).Where("id=?", 1).Update("hobby", "coding")
	// fmt.Printf("u:%#v\n", u)

	// 删除
	db.Delete(&u)
}
