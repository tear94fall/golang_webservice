package database

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id           int
	Title        string
	Content      string
	Writer       string
	RegisterDate string
	ModifiedDate string
	Views        int
}

func CreatePost(conn *gorm.DB, post *Post) error {
	err := conn.Create(post).Error
	if err != nil {
		return err
	}

	return nil
}

func GetPosts(conn *gorm.DB, post *[]Post) error {
	err := conn.Find(post).Error
	if err != nil {
		return err
	}

	return nil
}

func GetPostById(conn *gorm.DB, post *Post, id string) error {
	err := conn.Where("id = ?", id).First(post).Error
	if err != nil {
		return err
	}

	return nil
}

func GetPostRangeById(conn *gorm.DB, post *[]Post, start string, end string) error {
	err := conn.Where("id >= ? AND id <= ?", start, end).Find(&post).Error
	if err != nil {
		return err
	}

	return nil
}

func GetPostByWriter(conn *gorm.DB, post *Post, writer string) error {
	err := conn.Where("user_id = ?", writer).First(post).Error
	if err != nil {
		return err
	}

	return nil
}

func GetPostsCount(conn *gorm.DB, count *int64) error {
	err := conn.Model(&Post{}).Count(count).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdatePost(conn *gorm.DB, post *Post) error {
	err := conn.Save(post).Error
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(conn *gorm.DB, post *Post, id string) error {
	err := conn.Where("id = ?", id).Delete(post).Error
	if err != nil {
		return err
	}

	return nil
}
