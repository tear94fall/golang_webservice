package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "root"
	PASSWORD = "root1234"
	DATABASE = "my_db"
	HOST     = "docker.for.mac.localhost"
	PORT     = "3306"
)

type DBHandler struct {
	DBConn *gorm.DB
}

func NewDBHandler() (*DBHandler, error) {
	db_handler := &DBHandler{}

	db_info := USERNAME + ":" + PASSWORD + "@tcp" + "(" + HOST + ":" + PORT + ")/" + DATABASE + "?" + "parseTime=true&loc=Local"
	fmt.Println("db info : ", db_info)
	db, err := gorm.Open(mysql.Open(db_info), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil, err
	}

	connPool, err := db.DB()
	if err != nil {
		fmt.Printf("Error connection poll create : error=%v", err)
		return nil, err
	}

	connPool.SetMaxIdleConns(10)
	connPool.SetMaxOpenConns(100)
	connPool.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&Member{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Token{})

	db_handler.DBConn = db

	return db_handler, nil
}
