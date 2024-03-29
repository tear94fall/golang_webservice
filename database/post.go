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

func GetPostPaged(conn *gorm.DB, post *[]Post, current int, max int) error {
	var count int64 = 0

	offset := ((current - 1) * max)

	err := conn.
		Where("id > 0").
		Select("*").
		Offset(offset).
		Limit(max).
		Order(" id desc ").Find(&post).Count(&count).Error

	if err != nil {
		return nil
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
	// err := conn.Where("id = ?", id).Delete(post).Error
	err := conn.Unscoped().Delete(&Post{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
