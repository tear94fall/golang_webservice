package database

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Id       int
	UserId   string
	Password string
	Name     string
	Tel      string
}

func CreateMember(conn *gorm.DB, member *Member) error {
	err := conn.Create(member).Error
	if err != nil {
		return err
	}

	return nil
}

func GetMembers(conn *gorm.DB, member *[]Member) error {
	err := conn.Find(member).Error
	if err != nil {
		return err
	}

	return nil
}

func GetMemberById(conn *gorm.DB, member *Member, id string) error {
	err := conn.Where("id = ?", id).First(member).Error
	if err != nil {
		return err
	}

	return nil
}

func GetMemberByUserId(conn *gorm.DB, member *Member, userId string) error {
	err := conn.Where("user_id = ?", userId).First(member).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateMember(db *gorm.DB, member *Member) error {
	err := db.Save(member).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteMember(conn *gorm.DB, member *Member, id string) error {
	conn.Where("id = ?", id).Delete(member)
	return nil
}
