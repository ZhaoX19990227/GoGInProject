package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type UserInfo struct {
	// gorm.Model // 内嵌
	/**  包含以下四个字段
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	*/
	ID   int
	Name sql.NullString `gorm:"default:'小猪佩奇'"`
	//Name *string `gorm:"default:'小猪佩奇'"`
	//Name        string  `gorm:"default:'小猪佩奇'"`
	Age         int
	CreatedTime time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:zx123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	//禁止默认表名的复数形式
	db.SingularTable(true)
	defer db.Close()

	db.AutoMigrate(&UserInfo{})
	//
	//	userInfo := UserInfo{1,"小胖", 25, time.Now()}
	// name使用默认值插入到数据库
	// 注意：0，'',""，false，其他空值，都不会插入到数据库
	// 避免方式：使用指针或者实现Scanner/value的接口
	userInfo := UserInfo{Age: 55, Name: sql.NullString{String: "", Valid: true}, CreatedTime: time.Now()}

	//
	//db.Create(&wmChannel)
	// var userInfo UserInfo
	// 查询第一条记录
	//if err := db.First(&userInfo).Error;err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		fmt.Println("未找到匹配的记录")
	//	} else {
	//		fmt.Println("发生错误：" + err.Error())
	//	}
	//	return
	//}
	// 查询年龄大于18岁的用户
	//db.Find(&userInfo, "age > ?", 18)
	// 更新
	//db.Model(&userInfo).Update("Age", 20)
	// 删除
	//db.Delete(&userInfo)

	// 判断该对象是否为空
	fmt.Println(db.NewRecord(&userInfo))
	db.Debug().Create(&userInfo)
	fmt.Println(db.NewRecord(&userInfo))
	fmt.Printf("userInfo: %#v\n", userInfo)
}
