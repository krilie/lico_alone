package service

import (
	"bytes"
	"github.com/shurcooL/github_flavored_markdown"
	"github.com/yuin/goldmark"
	"io/ioutil"
	"os"
	"testing"
)

func TestToHtml(t *testing.T) {
	file, _ := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\bbb.txt")

	output := github_flavored_markdown.Markdown(file)
	//html := bluemonday.UGCPolicy().SanitizeBytes(output)

	err := ioutil.WriteFile("C:\\Users\\Administrator\\Desktop\\bbbc.html", output, os.ModePerm)
	println(err)
}

func TestDigProviderWithDao(t *testing.T) {
	file, _ := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\bbb.txt")
	var buf bytes.Buffer
	if err := goldmark.Convert(file, &buf); err != nil {
		panic(err)
	}
	err := ioutil.WriteFile("C:\\Users\\Administrator\\Desktop\\bbbc.html", buf.Bytes(), os.ModePerm)
	println(err)
}
