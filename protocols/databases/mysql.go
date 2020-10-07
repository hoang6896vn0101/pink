package databases

import (
	"fmt"
	configs "pink/settings/configs"

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
	conf := configs.MySQLConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName, conf.Password, conf.Host, conf.Port, conf.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
