package Services

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnect() (db *gorm.DB) {
	dsn := "root:password123@tcp(127.0.0.1:3306)/CytusDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return
}
