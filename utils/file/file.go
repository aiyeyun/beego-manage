package file

import (
	"mime/multipart"
	"errors"
	"strings"
	"path/filepath"
	"os"
	"log"
	"time"
	"crypto/md5"
	"io"
	"encoding/hex"
)

//图片上传
func ImagesUpload(file multipart.File, fileHeader *multipart.FileHeader, err error) (string, error) {
	//defer file.Close()
	suffix := make(map[string]string)
	suffix["png"] = "png"
	suffix["jpg"] = "jpg"
	suffix["gif"] = "gif"
	if err != nil {
		return "", errors.New("未上传文件")
	}

	ok := validateSuffix(fileHeader.Filename, suffix)
	if !ok {
		return "", errors.New("文件不合法")
	}

	//保存文件
	return saveFile(file, fileHeader)
}

//检验后缀是否合法
func validateSuffix(fileName string, suffixs map[string]string) (bool) {
	ary := strings.Split(fileName, ".")
	suffix := ary[len(ary) - 1]
	if _, ok := suffixs[suffix]; !ok {
		return false
	}
	return true
}

//保存文件
func saveFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	//检查文件夹是否存在
	path := getCurrentDirectory()
	tofile := "/static/upload/image/" + time.Now().Format("2006-01-02")
	saveFilePath := tofile
	tofile = path + tofile
	if b, _ := pathExists(tofile); !b {
		os.Mkdir(tofile, 0777)
	}
	tempFile, _ := fileHeader.Open()
	md5h := md5.New()
	io.Copy(md5h, tempFile)
	tofile += "/" + hex.EncodeToString(md5h.Sum(nil)) + "." + GetFileSuffix(fileHeader.Filename)
	//tofile = tofile + "/" + hex.EncodeToString(md5h.Sum(nil)) + "." + GetFileSuffix(fileHeader.Filename)
	os.Create(tofile)
	f, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		return "", err
	}
	io.Copy(f, file)
	return saveFilePath + "/" + hex.EncodeToString(md5h.Sum(nil)) + "." + GetFileSuffix(fileHeader.Filename), nil
}

//获取当前目录
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//文件 或 文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//获取文件后缀
func GetFileSuffix(fileName string) string {
	ary := strings.Split(fileName, ".")
	return ary[len(ary) - 1]
}