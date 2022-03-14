package attach

import (
	"errors"
	"fmt"
	"main/common"
	"main/database"
	"main/util"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AttachInfo struct {
	ArticleId string
	FileName  string
	FileSize  string
}

func Upload(c *gin.Context) {

}

func SaveAttachFile(c *gin.Context, id string) error {
	file, _ := c.FormFile("file")

	if file == nil {
		err := errors.New("첨부 파일이 없습니다")
		return err
	}

	file_name := filepath.Base(file.Filename)
	file_size := strconv.FormatInt(file.Size, 10)
	save_dir, _ := util.EncStr(id)
	save_path := common.AttachPath + "/" + save_dir
	full_file_path := save_path + "/" + file_name

	if !util.CheckExistFile(save_path) {
		util.MakeDirectory(save_path)
	}

	if err := c.SaveUploadedFile(file, full_file_path); err != nil {
		return err
	}

	attach := &AttachInfo{id, file_name, file_size}
	upload := NewAttach(attach)

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.CreateAttach(db.DBConn, upload)

	return nil
}

func NewAttach(attach *AttachInfo) *database.Attach {
	upload := &database.Attach{}
	upload.ArticleId = attach.ArticleId
	upload.FileName = attach.FileName
	upload.FileSize = attach.FileSize

	return upload
}

func UpdateAttachFile(c *gin.Context, file *multipart.FileHeader, old_file string, id int) error {
	/*
		1. 파일 변경이 없는 경우
		2. 파일이 추가된 경우
		3. 파일이 삭제된 경우
		4. 파일이 변경된 경우
	*/

	file_name := ""
	if file != nil {
		file_name = filepath.Base(file.Filename)
	}

	file_dir, _ := util.EncStr(strconv.Itoa(id))
	file_path := common.AttachPath + "/" + file_dir
	full_file_path := file_path + "/" + file_name

	if (old_file == "" && file_name == "") || (old_file == file_name) {
		return nil
	}

	if old_file == "" && file_name != "" {
		util.MakeDirectory(file_path)
		if err := c.SaveUploadedFile(file, full_file_path); err != nil {
			return err
		}
	}

	if old_file != "" {
		if file_name == "" {
			fmt.Println(file_path)
			util.RemoveDriectory(file_path)
		} else if file_name != "" {
			util.RemoveFile(file_path + "/" + old_file)
			if err := c.SaveUploadedFile(file, full_file_path); err != nil {
				return err
			}
		}
	}

	return nil
}
