package files

import (
	"os"
	"path"
)

// SplitNameExt splits the name and the extension name.
// The filename is cosidered as follows: 'name.name.name.ext'.
func SplitNameExt(filename string) (string, string) {
	/*name.name.name.ext*/
	if filename == "" {
		return "", ""
	}
	if filename[0] == '.' {
		name, ext := SplitNameExt(filename[1:])
		return "." + name, ext
	}
	ext := path.Ext(filename)
	if ext == "" {
		return filename, ""
	}
	return filename[:len(filename)-len(ext)], ext[1:]
}

// JoinNameExt return the name.ext of the file.
func JoinNameExt(name, ext string) string {
	if ext == "" {
		return name
	}
	return name + "." + ext
}

// CreateDir crea el directorio si no existe
func CreateDir(path string) error {
	return os.MkdirAll(path, 0764)
}

// CreateDirPanic crea el directorio si no existe
func CreateDirPanic(path string) {
	if err := os.MkdirAll(path, 0764); err != nil {
		panic("Cant create directory " + path + ": " + err.Error())
	}
}
