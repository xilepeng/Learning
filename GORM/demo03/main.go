package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// type Model struct {
// 	ID        uint `gorm:"primarykey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt DeletedAt `gorm:"index"`
// }

// 表名默认就是结构体名称的复数，例如：
type User struct { // 默认表名是 `users`
	gorm.Model // 内嵌 gorm.Model

	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

type UserDefine struct {
	ID   uint // 名为`ID`的字段会默认作为表的主键
	Name string
}

type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
}

// 将 User 的表名设置为 `profiles`
func (User) TableName() string {
	return "mojoman"
}

// func (u User) TableName() string {
// 	if u.Role == "admin" {
// 		return "admin_users"
// 	} else {
// 		return "users"
// 	}
// }

// create  database  test character set utf8mb4 collate utf8mb4_unicode_ci;
// use test;
func main() {
	dsn := "root:Yizhili80@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         nil,
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 单数表名
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &UserDefine{}, &Animal{})
	db.Table("user").Take(&User{})

}
