package utils

import (
	"errors"
	"github.com/astaxie/beego"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"
)

// @title 多图片上传
func UploadImages(files []*multipart.FileHeader, catePath string) (fileArr []string, err error) {
	if len(files) == 0 {
		errors.New("传入的文件夹为空")
		return
	}
	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		fileFullName, sErr := UploadImage(files[i], catePath)
		if sErr != nil {
			continue
		}
		fileArr = append(fileArr, fileFullName)
	}
	if len(fileArr) == 0 {
		errors.New("文件上传失败")
	}
	return
}

// @title单图片上传
func UploadImage(file *multipart.FileHeader, catePath string) (fileFullName string, err error) {
	f, err := file.Open()
	defer f.Close()
	if err != nil {
		errors.New("文件打开失败")
		return
	}
	basePath := beego.AppConfig.String("file_upload_path")
	headPath := beego.AppConfig.String("file_head")
	datePath := time.Now().Format("2006/01/02")
	fileSuffix := path.Ext(file.Filename)
	filePath := basePath + catePath + "/" + datePath + "/"
	if _, err = os.Stat(filePath); err != nil {
		err = os.MkdirAll(filePath, 0755)
		if err != nil {
			errors.New("文件夹权限不足")
		}
	}
	fileName := GetRandomString(15, false) + fileSuffix
	createFile := filePath + fileName
	fileFullName = headPath + catePath + "/" + datePath + "/" + fileName
	dst, err := os.Create(createFile)
	defer dst.Close()
	if err != nil {
		errors.New("创建文件失败")
	}
	_, err = io.Copy(dst, f)
	if err != nil {
		errors.New("文件复制失败")
	}
	return
}
