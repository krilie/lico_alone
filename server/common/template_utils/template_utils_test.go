package template_utils

import (
	"os"
	"testing"
	"text/template"
)

func TestGetProjectPath(t *testing.T) {
	println(GetProjectPath())
}

func TestGenDir(t *testing.T) {
	GenDir("common\\gorm_template", "module\\module-test",
		map[string]interface{}{
			"ParentPackage": "module-test",
			"ModuleName":    "TestArticle",
			"models":        []string{"One", "Two"},
		})
}

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func TestGenDir2(t2 *testing.T) {
	f1 := Friend{Fname: "xiaofang"}
	f2 := Friend{Fname: "wugui"}
	t := template.New("test")
	t = template.Must(t.Parse(
		`hello {{.UserName}}!
{{ range .Emails }}
an email {{ . }}
{{- end }}
{{ with .Friends }}
{{- range . }}
my friend name is {{.Fname}}
{{- end }}
{{ end }}`))
	p := Person{UserName: "longshuai",
		Emails:  []string{"a1@qq.com", "a2@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
