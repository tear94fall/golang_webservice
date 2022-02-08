package database

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Id           int
	ArticleId    string
	Writer       string
	Content      string
	RegisterDate string
}

func CreateComment(conn *gorm.DB, comment *Comment) error {
	err := conn.Create(comment).Error
	if err != nil {
		return err
	}

	return nil
}

func GetCommentByArticleId(conn *gorm.DB, comment *[]Comment, id string) error {
	err := conn.Where("article_id = ?", id).Find(comment).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteCommentByCommentId(conn *gorm.DB, id string) error {

	return nil
}

func DeleteCommentByArticleId(conn *gorm.DB, id string) error {

	return nil
}

func DeleteComment(conn *gorm.DB, comment *Comment) error {

	return nil
}

func UpdateComment(conn *gorm.DB, token string, id string) error {
	t := &Token{}

	if err := GetTokenById(conn, t, id); err != nil {
		return err
	}

	t.Login = token

	err := conn.Save(t).Error
	if err != nil {
		return err
	}

	return nil
}
