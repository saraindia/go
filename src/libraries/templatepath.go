//Gets the template path
package libraries

import (
	"os"
	"path/filepath"
)

//Gets the template path
func GetTemplatePath() string {
	//var templatePath string = "./src/templates"
	dir, _ := os.Getwd()
	return filepath.Join(dir)
}
