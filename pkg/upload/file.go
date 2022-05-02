package upload

import (
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
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
	fileName = util.Md5(fileName)
	return fileName + fileExt
}

func GetFileExts(str string) string {
	return path.Ext(str) //获取文件后缀
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

//检查文件或者目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

//判断当前文件后缀类型
func CheckContainExt(f FileType, name string) bool {
	ext := GetFileExts(name)
	ext = strings.ToUpper(ext)
	switch f {
	case TypeImage:
		s := global.AppSetting.UploadImageAllowExts
		for _, v := range s {
			gExt := strings.ToUpper(v)
			if ext == gExt {
				return true
			}
		}
	}
	return false
}

func CheckFileSize(f FileType, fileMult multipart.File) bool {
	b, _ := ioutil.ReadAll(fileMult)
	size := len(b)
	switch f {
	case TypeImage:
		if size > global.AppSetting.UploadImageMaxSize*1024*1024 {
			return false
		}
	}
	return true
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

//创建文件保存目录
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, f)
	return err
}
