/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/22
 * Time 10:17 上午
 */

package dao

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"

	"gorm.io/driver/postgres"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"imooc.com/ccmouse/learngo/gin-gorm/config"
)

var db *gorm.DB

func setup() {
	var dbURI string
	var dialector gorm.Dialector
	if config.DatabaseSetting.Type == "mysql" {
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&&parseTime=true",
			config.DatabaseSetting.User,
			config.DatabaseSetting.Password,
			config.DatabaseSetting.Host,
			config.DatabaseSetting.Port,
			config.DatabaseSetting.Name,
		)
		dialector = mysql.New(mysql.Config{
			DSN:                       dbURI,
			SkipInitializeWithVersion: false,
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
		})
	} else if config.DatabaseSetting.Type == "postgres" {
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
			config.DatabaseSetting.Host,
			config.DatabaseSetting.Port,
			config.DatabaseSetting.User,
			config.DatabaseSetting.Name,
			config.DatabaseSetting.Password,
		)
		dialector = postgres.New(postgres.Config{
			DSN:                  dbURI,
			PreferSimpleProtocol: true,
		})
	} else {
		dbURI = fmt.Sprintf("bms.db")
		dialector = sqlite.Open("bms.db")
	}
	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Print(err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		fmt.Errorf("connect db server failed. %s", err.Error())
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Second * 600)
	db = conn
}

func GetDB() *gorm.DB {
	if db == nil {
		setup()
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Errorf("connect db server failed. %s", err.Error())
		setup()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		setup()
	}
	return db
}
