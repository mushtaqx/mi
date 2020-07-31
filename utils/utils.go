package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"unicode"
)

// SnakeCase : converts PascalCase to snake_case
func SnakeCase(pascalCase string) (res string) {
	if pascalCase == "" {
		return ""
	}

	for i := range pascalCase {
		// Store character to new result var
		res += string(pascalCase[i])
		if unicode.IsUpper(rune(pascalCase[i])) {
			// Convert to lowercase
			tolower := strings.ToLower(string(pascalCase[i]))
			// Remove last added uppercase char from result string.
			res = strings.TrimSuffix(res, string(pascalCase[i]))
			// Add an underscore before lowercased char.
			if i != 0 {
				res += "_"
			}
			// Add lowercased char to result string.
			res += tolower
		}
	}

	return res
}

// FileExist : checks if any file in dir listing has given name
func FileExist(p, name string) os.FileInfo {
	// Get dir entries
	info, err := ioutil.ReadDir(path.Clean(p))
	if err != nil {
		log.Fatal("Couldn't read directory for listing")
	}

	for _, i := range info {
		if strings.Contains(i.Name(), name) && !i.IsDir() {
			return i
		}
	}

	return nil
}
