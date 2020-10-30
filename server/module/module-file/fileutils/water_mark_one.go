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
	imgbytes, err := base64.StdEncoding.DecodeString(markImg)
	if err != nil {
		panic(err)
	}
	markb = imgbytes
}

var markImg = "R0lGODlhfQAUAPcAAAAAAAEBAQICAgMDAwQEBAUFBQYGBgcHBwgICAkJCQoKCgsLCwwMDA0NDQ4ODg8PDxAQEBERERISEhMTExQUFBUVFRYWFhcXFxgYGBkZGRoaGhsbGxwcHB0dHR4eHh8fHyAgICEhISIiIiMjIyQkJCUlJSYmJicnJygoKCkpKSoqKisrKywsLC0tLS4uLi8vLzAwMDExMTIyMjMzMzQ0NDU1NTY2Njc3Nzg4ODk5OTo6Ojs7Ozw8PD09PT4+Pj8/P0BAQEFBQUJCQkNDQ0REREVFRUZGRkdHR0hISElJSUpKSktLS0xMTE1NTU5OTk9PT1BQUFFRUVJSUlNTU1RUVFVVVVZWVldXV1hYWFlZWVpaWltbW1xcXF1dXV5eXl9fX2BgYGFhYWJiYmNjY2RkZGVlZWZmZmdnZ2hoaGlpaWpqamtra2xsbG1tbW5ubm9vb3BwcHFxcXJycnNzc3R0dHV1dXZ2dnd3d3h4eHl5eXp6ent7e3x8fH19fX5+fn9/f4CAgIGBgYKCgoODg4SEhIWFhYaGhoeHh4iIiImJiYqKiouLi4yMjI2NjY6Ojo+Pj5CQkJGRkZKSkpOTk5SUlJWVlZaWlpeXl5iYmJmZmZqampubm5ycnJ2dnZ6enp+fn6CgoKGhoaKioqOjo6SkpKWlpaurprGxqMC/q83LreDer/bysv35sv76sv76sv76sv76sv76sv76s/76s/76s/76s/76s/76s/77s/77tv77u/77w/780v795v7+9v7+/f7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v7+/v////7+/v7+/v7+/v///yH/C05FVFNDQVBFMi4wAwEAAAAh+QQJAQD7ACwAAAAAfQAUAAAI/gD3CRxIsKDBgwgTKlzIsKHDhxAjSpxIsaLFixgzatxoEBeugx5Dihw5UqDHgidJhuTI0qHKkvtOmvw4M+bHlzJTlpTZsmfDnDRtDuSpc2VNgkVJ+lwKkibQoUGJ3pyKE+ZPo1CZTqz69KjQpFxfKuTZkixGpUjDqq0qNKFZjm8r6rQZtW5WqGtVjq0rVeTMuH+N+kXbdC7etojFFtYL2KtXw10RSq0Z2S3hxHbB4gyMtSnKoI+dgu5o96vohWRTl6a7szNrx5Y/yz7aGDPl05Lnmp18m3XpsLHTCodMlfRd4i7f5kWrGTbpwa39vqZbODDnvdBnG896WbtWi7W/EosfP528+fPo06tfz578v38BAQA7"
var markb []byte

// 添加水印one 简单的
func WaterMarkOne(ctx context.Context, oriImg io.Reader, outImg io.Writer) error {
	// oriImag 原始图像
	img, err := jpeg.Decode(oriImg)
	if err != nil {
		return err
	}
	// 水印图像
	markImg, err := gif.Decode(bytes.NewReader(markb))
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
