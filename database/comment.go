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

func GetCommentById(conn *gorm.DB, comment *Comment, id string) error {
	err := conn.Where("id = ?", id).First(comment).Error
	if err != nil {
		return err
	}

	return nil
}

func GetArticleIdByCommentId(conn *gorm.DB, comment_id string) (string, error) {
	comment := &Comment{}
	err := conn.Where("id = ?", comment_id).First(comment).Error
	if err != nil {
		return "", err
	}

	return comment.ArticleId, nil
}

func DeleteCommentByCommentId(conn *gorm.DB, id string) error {

	return nil
}

func DeleteCommentByArticleId(conn *gorm.DB, id string) error {

	return nil
}

func DeleteComment(conn *gorm.DB, comment *Comment) error {
	err := conn.Unscoped().Delete(&Comment{}, comment.Id).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateComment(conn *gorm.DB, comment *Comment) error {
	err := conn.Save(comment).Error
	if err != nil {
		return err
	}

	return nil
}
