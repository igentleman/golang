package service

import (
	"fmt"
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (s *Service) UploadFile(FileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(FileType, fileName) {
		return nil, fmt.Errorf("文件后缀错误")
	}
	if !upload.CheckFileSize(FileType, file) {
		return nil, fmt.Errorf("文件太大")
	}
	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("创建文件失败")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, fmt.Errorf("文件权限不足")
	}
	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
