package template_utils

import "testing"

// TestGenModuleCatchword 流行语模块
func TestGenModuleCatchword(t *testing.T) {
	GenDir("common\\gorm_template", "module\\module-union",
		map[string]interface{}{
			"packageName": "module-union",
			"moduleName":  "Union",
			"models":      []string{"Union"},
		})
}
