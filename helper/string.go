package helper

import (
	"path/filepath"
	"runtime"
	"strings"
)

func SanitizedName(filename string) string {
	if len(filename) > 1 && filename[1] == ':' &&
		runtime.GOOS == "windows" {
		filename = filename[2:]
	}

	filename = filepath.ToSlash(filename)
	filename = strings.TrimLeft(filename, "/.")

	return strings.Replace(filename, "../", "", -1)
}
