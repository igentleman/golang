package upload

import (
	"goproject/main/ginweb/global"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage FileType = iota + 1
)

func GetFileName(name string) string {
	fileExt := GetFileExts(name)
	fileName := strings.TrimSuffix(name, fileExt)
	return fileName + fileExt
}

func GetFileExts(str string) string {
	return path.Ext(str)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}
