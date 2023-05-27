package configs

// Here we will make our connection to DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB //Making a variable Db pointing towards gorm.db
)

// user:password@/dbname?charset=utf8&parseTime=True&loc=Local

func ConnectToDb() {
	d, err := gorm.Open("mysql", "rounak:Rounak@00@/LocalDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB has a return type gorm.DB
func GetDB() *gorm.DB {
	return db
}
