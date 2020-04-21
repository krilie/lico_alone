package file_util

import (
	"io"
	"net/http"
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
