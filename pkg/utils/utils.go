package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// SnakeCase : converts str to snake_case
func SnakeCase(str string) (res string) {
	if str == "" {
		return ""
	}

	for i := range str {
		// Store character to new result var
		if !unicode.IsUpper(rune(str[i])) {
			res += string(str[i])
			continue
		}
		// Add an underscore before lowercase char.
		if i > 0 {
			res += "_"
		}
		// Add lowercased char to result string.
		res += strings.ToLower(string(str[i]))
	}

	return res
}

// ToUpperFirst : uppercase first character of str
func ToUpperFirst(str string) string {
	if unicode.IsUpper(rune(str[0])) {
		return str
	}
	// transform lowercase str to uppercase
	upper := strings.ToUpper(string(str[0]))
	upper += str[1:]

	return upper
}

// FileNameLike : Gets first result containing name
// in filename of dir listing as full or substring
func FileNameLike(name, p string) (file os.FileInfo, err error) {
	// Get dir entries
	files, err := ioutil.ReadDir(path.Clean(p))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if strings.Contains(f.Name(), name) && !f.IsDir() {
			return f, nil
		}
	}

	return nil, fmt.Errorf("file %s not found in %s", name, p)
}

// NowSpecial : get full time.now() formated as int
func NowSpecial() (t int) {
	now := time.Now().Format("020106030405.000")
	if t, err := strconv.Atoi(strings.Replace(now, ".", "", 1)); err == nil {
		return t
	}

	return t
}
