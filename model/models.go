package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zou2699/learnGin2/utils/setting"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var err error
	db, err = gorm.Open(setting.DBConfig.Dialect, setting.DBConfig.URL)
	if err != nil {
		log.Panic("Failed to connect to DB:", err.Error())
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DBConfig.TablePrefix + defaultTableName
	}

	// 让grom转义struct名字的时候不用加上s
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
