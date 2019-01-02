package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zou2699/learnGin2/pkg/setting"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        int `form:"id" gorm:"primary_key" json:"id" bind:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
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

	//debug
	db.LogMode(true)
}

func CloseDB() {
	defer db.Close()
}
