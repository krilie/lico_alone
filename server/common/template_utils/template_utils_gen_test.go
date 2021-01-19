package template_utils

import "testing"

// TestGenModuleCatchword 流行语模块
func TestGenModuleCatchword(t *testing.T) {
	GenDir("common\\gorm_template", "module\\module-catchword",
		map[string]interface{}{
			"packageName": "module-catchword",
			"moduleName":  "Catchword",
			"models":      []string{"Catchword"},
		})
}
