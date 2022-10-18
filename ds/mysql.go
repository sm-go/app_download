package ds

import (
	"app-download/config"
	"app-download/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMysql() (*gorm.DB, error) {
	host := config.MysqlDB.Host
	port := config.MysqlDB.Port
	user := config.MysqlDB.User
	pass := config.MysqlDB.Pass
	dbname := config.MysqlDB.DB

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error on connection to database")
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.App{},
		&model.Domain{},
		&model.DownloadLog{},
		&model.InstallLog{},
	); err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
