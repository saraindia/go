package libraries

import (
	"testing"
	"os"
)

func TestTemplatpath(t *testing.T)  {
	var pathstr = GetTemplatePath()
	dir, _ := os.Getwd()
	if pathstr != dir {
		t.Error("The path is wrong")
	}
}