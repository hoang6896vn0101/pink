package databases

import (
	"fmt"
	"pink/config/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL struct
type MySQL struct{}

// DB func
// Arguments:
// - No
// Return:
// 1. *gorm.DB
func DB() *gorm.DB {
	conf := env.MySQLConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName, conf.Password, conf.Host, conf.Port, conf.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error: Connection fails")
	}
	return db
}
