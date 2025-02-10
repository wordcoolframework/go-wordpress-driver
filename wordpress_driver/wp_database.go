package wordpressdriver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() error {
	var err error

	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/wordpress"), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	return nil
}
