package module

import (
	"log"
	"os"
	"path"
	"reflect"
	"strings"
)

func FindPackagesForType(any interface{}) string {
	return reflect.TypeOf(any).PkgPath()
}

// getProjectInformation will return the module name and the path to the directory
// We need the full file path, we also assume the path is in the format of the current OS
// getProjectInformation("C:\Code\fiber-code-gen\example\basic") will return "code-gen", "example\basic", nil
// Because go.mod is in C:\Code\fiber-code-gen and the module name is code-gen
// getProjectInformation("C:\Code\fiber-code-gen") will return "code-gen", "", nil
func getProjectInformation(filePath, startingPoint string) (string, string, error) {
	filePath = path.Clean(filePath)
	if startingPoint == "" {
		startingPoint = filePath
	}
	goModBytes, err := os.ReadFile(filePath + string(os.PathSeparator) + "go.mod")
	if err != nil {
		splitPath := strings.Split(filePath, string(os.PathSeparator))
		if len(splitPath) <= 1 {
			log.Fatal("go.mod not found in directories `" + filePath + "` to `" + startingPoint + "`")
		}
		return getProjectInformation(strings.Join(splitPath[:len(splitPath)-1], string(os.PathSeparator)), startingPoint)
	}

	lines := strings.Split(string(goModBytes), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module") {
			modulePath := strings.TrimSpace(strings.TrimPrefix(line, "module"))
			return modulePath, strings.Replace(strings.Replace(startingPoint, filePath+string(os.PathSeparator), "", -1), "\\", "/", -1), nil
		}
	}

	return "", "", nil
}
