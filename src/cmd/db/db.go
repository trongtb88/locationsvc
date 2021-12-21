package db

import (
	"fmt"
	"github.com/trongtb88/locationsvc/src/business/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)


func ConnectDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
	var (
		db *gorm.DB
		err error
	)
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		db, err = gorm.Open(mysql.Open(DBURL))
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	err = db.Debug().AutoMigrate(&entity.Place{}) //database migration

	if err != nil {
		log.Fatal("Error when migration table:", err)
	}
	return db
}

