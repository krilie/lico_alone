package template_utils

import "testing"

func TestGetProjectPath(t *testing.T) {
	println(GetProjectPath())
}

func TestGenDir(t *testing.T) {
	GenDir("common\\gorm_template", "module\\module-test",
		map[string]interface{}{
			"ParentPackage": "ParentPackage",
			"ModuleName":    "ModuleName",
			"models":        []string{"aa", "bb"},
		})
}
