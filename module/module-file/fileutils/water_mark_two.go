package fileutils

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"io"
)

func init() {
	var img = "R0lGODlhZABkAPcAAHZ2dnR0dHx8fIiIiJaWlqWlpaysrMDAwNHR0d7e3urq6vT09Pz8/P39/f7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v////7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v////7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v////7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v////7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v////7+/v7+/v7+/v///yH/C05FVFNDQVBFMi4wAwEAAAAh+QQJAQDSACwAAAAAZABkAAAI/gClCRxIsKDBgwgTKlzIsKHDhxAjSpxIsaLFixgzatzIsaPHjyBDihxJsqTJkyhTqlzJsqXLlzBjypxJs6bNmzhz6tzJs6fPn0CDCh1KtKjRo0iTKl1qMcECphQNCAggoADUiAaoDggQYMBVhwsEEBCYYOrYrwoLTD0wsGwABWgTFuBqgOBctnERzhWQYGBWvHkFLnAqcO+BBAgECAg8kMBUsXDncqUKOO/WAQMeR6b8lPGBAHgZbD27ty/jraalhRXQeS/cwKgJxi4s4HXerF4HZu4sGCpgAqYZTB2QIMHns2gJBLC6FbDbybXzul1bkIEBzJXjukXOOOHWqVYH99ruLu2u2/Dlc5OXhoAseGnH1xucTpW3fPGOC9i/z7+///8ABvjRAdgZ1N56W0GXWnPkKXdYAt8NpFZqeYWlnjQM9tZdAl0RhECH9ylAFQNtgdjdgRh2RSJ8oJG31YEKPIYZVS5yBZgCCQZAwHhxNSdAdgokwGOPLf7432d4CVCXQCh+lR1ZAwmQ22foQbVVlQZJyeKFvrXo3QCfcXlVZ08KlJmJgc2FpZloBvbZcvNx192ba35lnEJ0BiaclwjlmdebZQ705pJxAYonX57ZqNCKifIJoKECQhqgpI9y1eSRYgqo6aacdurpp6CGKuqopJaK0j//BAQAOw=="
	imgbytes, err := base64.StdEncoding.DecodeString(img)
	if err != nil {
		panic(err)
	}
	markbLizo = imgbytes
}

var markbLizo []byte

// 添加水印one 简单的
func WaterMarkTwo(ctx context.Context, oriImg io.Reader, outImg io.Writer) error {
	// oriImag 原始图像
	img, err := jpeg.Decode(oriImg)
	if err != nil {
		return err
	}
	// 水印图像
	markImg, err := gif.Decode(bytes.NewReader(markbLizo))
	if err != nil {
		return err
	}
	offset := image.Pt(img.Bounds().Dx()-markImg.Bounds().Dx()-10, img.Bounds().Dy()-markImg.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, img, image.Point{}, draw.Src)
	draw.Draw(m, markImg.Bounds().Add(offset), markImg, image.Point{}, draw.Over)
	err = jpeg.Encode(outImg, m, &jpeg.Options{100})
	return err
}
