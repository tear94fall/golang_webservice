package database

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Id    int
	Login string
}

func CreateToken(conn *gorm.DB, token string) error {

	t := &Token{}
	t.Login = token

	err := conn.Create(t).Error
	if err != nil {
		return err
	}

	return nil
}

func CheckToken(conn *gorm.DB, token string) (int64, error) {
	var count int64

	err := conn.Model(&Token{}).Where("login = ?", token).Count(&count)
	if err != nil {
		return 0, err.Error
	}

	return count, nil
}
