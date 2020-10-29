// +build !auto_test

package service

import (
	"github.com/shurcooL/github_flavored_markdown"
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
