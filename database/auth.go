package database

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Id     int
	Login  string
	UserId string
}

func CreateToken(conn *gorm.DB, token string, id string) error {
	t := &Token{}
	t.Login = token
	t.UserId = id

	err := conn.Create(t).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteToken(conn *gorm.DB, token string, id string) error {
	t := &Token{}

	if err := GetIdByToken(conn, t, token); err != nil {
		return err
	}

	err := conn.Where(&Token{Login: token, UserId: id}).Delete(t).Error
	if err != nil {
		return nil
	}

	return nil
}

func GetIdByToken(conn *gorm.DB, token *Token, login_token string) error {
	err := conn.Where("login = ?", login_token).First(token).Error
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
