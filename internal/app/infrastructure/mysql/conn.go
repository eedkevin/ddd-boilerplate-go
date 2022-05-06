package mysql

import (
	"fmt"
	"log"

	"ddd-boilerplate/internal/app/adapter/cfg"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Client is the alias of gorm.DB
type Client = gorm.DB

//Connection creates a mysql client
func Connection(cfg cfg.MysqlConfig) *Client {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Schema)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error on connecting to mysql db", err)
	}
	return db
}
