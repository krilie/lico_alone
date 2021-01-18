package template_utils

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenFiles(origin, out string, vars interface{}) {
	tmpl := template.New("genCode")
	originFile, err := os.Open(origin)
	if err != nil {
		panic(err)
	}
	defer originFile.Close()
	content, err := ioutil.ReadAll(originFile)
	if err != nil {
		panic(err)
	}
	tmpl, err = tmpl.Parse(string(content))
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	var buf = &bytes.Buffer{}
	if err := tmpl.Execute(buf, vars); nil != err {
		panic(err)
	}
	source, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	write, err := outFile.Write(source)
	if err != nil || write != len(source) {
		panic(err)
	}
	println(out)
}

// dir out 都是相对路径的目录
func GenDir(inDir, outDir string, vars interface{}) {

	currentPath := GetProjectPath()
	err := filepath.Walk(currentPath+"\\"+inDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			outDir, outPath := BuildGoFilePath(currentPath, path, currentPath+"\\"+inDir, info.Name(), outDir)
			if _, err := os.Stat(outDir); os.IsNotExist(err) {
				err := os.MkdirAll(outDir, os.ModePerm)
				if err != nil {
					panic(err)
				}
			}
			GenFiles(path, outPath, vars)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func GetProjectPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	index := strings.LastIndex(currentPath, `\server\`)
	return currentPath[:index+7]
}

// C:\project\lico_alone\server = projectPath
// module\module-time-oc = outPath
// C:\project\lico_alone\server\module\module-time-oc = outFullPath
// oriFileBasePath = C:\project\lico_alone\server\common\gorm_template
func BuildGoFilePath(projectPath, oriFilePath, oriFileBasePath, oriFileName, outPath string) (dirPath, outFilePath string) {
	// C:\project\lico_alone\server\common\gorm_template\dao\dao.gohtml => \dao\dao.gohtml
	var oriFileRelativePath = strings.ReplaceAll(oriFilePath, oriFileBasePath, "")
	// \dao\
	var midPath = strings.ReplaceAll(oriFileRelativePath, oriFileName, "")
	// C:\project\lico_alone\server\module\module-time-oc\dao\
	var outDir = projectPath + "\\" + outPath + midPath
	var outFileName = strings.ReplaceAll(oriFileName, ".gohtml", ".go")
	return outDir, outDir + outFileName
}
