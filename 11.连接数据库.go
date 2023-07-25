package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type UserInfo struct {
	ID          int
	Name        string
	Age         int
	CreatedTime time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:zx123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	//db.AutoMigrate(&UserInfo{})
	//
	//wmChannel := UserInfo{1, "小胖", 25, time.Now()}
	//
	//db.Create(&wmChannel)
	var userInfo UserInfo
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
	fmt.Printf("userInfo: %#v\n", userInfo)
}
