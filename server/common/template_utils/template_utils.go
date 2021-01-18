package template_utils

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenFiles(tmpl *template.Template, origin, out string, vars interface{}) {
	tmpl, err := tmpl.ParseFiles(origin)
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	if err := tmpl.Execute(outFile, vars); nil != err {
		panic(err)
	}
}

// dir out 都是相对路径的目录
func GenDir(inDir, outDir string, vars interface{}) {
	tmpl := template.New("gen code")
	//tmpl.Funcs()

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
			GenFiles(tmpl, path, outPath, vars)
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
