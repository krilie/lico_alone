package file_util

import (
	"io"
	"net/http"
	"strings"
)

func GetContentType(file io.ReadSeeker) (string, error) {
	decByte := make([]byte, 512)
	if _, err := file.Read(decByte); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}
	contentType := http.DetectContentType(decByte)
	return contentType, nil
}

func GetFileExtension(fileName string) string {
	index := strings.LastIndex(fileName, ".")
	if index == -1 {
		return ""
	}
	return fileName[index:]
}
