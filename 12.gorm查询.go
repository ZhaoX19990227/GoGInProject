package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
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

	db.AutoMigrate(&User{})

	//for i := 0; i < 10; i++ {
	//	user1 := User{Name: "小胖" + strconv.Itoa(i), Age: int64(i)}
	//	db.Create(&user1)
	//}

	user := new(User)
	//// 查找第一条记录
	//if err := db.First(&user).Error;err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		fmt.Println("未找到匹配的记录")
	//	} else {
	//		fmt.Println("发生错误：" + err.Error())
	//	}
	//	return
	//}
	//fmt.Printf("userInfo: %#v\n", user)
	//
	//// 查询所有
	// var users []User
	//db.Debug().Find(&users)
	//fmt.Printf("users: %#v\n", users)
	//
	// where查询 一条记录
	//db.Where("name = ?","小胖9").First(&user)
	//fmt.Printf("userInfo: %#v\n", user)
	// where查询 多条记录
	//db.Where("name like ?","小胖%").Find(&users)
	//fmt.Printf("users: %#v\n", users)
	// <>
	//db.Debug().Where("name <> ?","大胖").Find(&users)
	//fmt.Printf("users: %#v\n", users)
	// in
	//db.Debug().Where("age in (?)",[]int{1,2,3}).Find(&users)
	//fmt.Printf("users: %#v\n", users)

	// struct
	//db.Debug().Where(&User{Name: "小胖0",Age: 0}).First(&user)
	//fmt.Printf("user: %#v\n", user)

	// map
	//db.Debug().Where(map[string]interface{}{"name":"小胖1","age":"1"}).Find(&user)
	//fmt.Printf("user: %#v\n", user)
	// 主键的切片
	//db.Debug().Where([]int64{5,6,7}).Find(&users)
	//fmt.Printf("users: %#v\n", users)

	//FirstOrInit  没有就init一个对象
	//db.FirstOrInit(&user,User{Name: "小可爱"})
	//fmt.Printf("user: %#v\n", user)

	//查不到的话多个字段初始化
	db.Attrs(User{Age: 101}).FirstOrInit(&user, User{Name: "小可爱"})
	fmt.Printf("user: %#v\n", user)
}
