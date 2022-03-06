package database

import (
	"gorm.io/gorm"
)

type Attach struct {
	gorm.Model
	Id        int
	ArticleId string
	FileName  string
	FileSize  string
}

func CreateAttach(conn *gorm.DB, attach *Attach) error {
	err := conn.Create(attach).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAttachById(conn *gorm.DB, attach *Attach, id string) error {
	err := conn.Where("id = ?", id).First(attach).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAttachByArticleId(conn *gorm.DB, attach *[]Attach, id string) error {
	err := conn.Where("article_id = ?", id).Find(attach).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteAttach(conn *gorm.DB, attach *Attach) error {
	err := conn.Unscoped().Delete(&Attach{}, attach.Id).Error
	if err != nil {
		return err
	}

	return nil
}
