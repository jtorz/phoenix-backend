package files

import (
	"os"
	"path"
)

// SplitNameExt splits the name and the extension from the filename.
// The filename is cosidered as follows: 'name.name.name.ext'.
func SplitNameExt(filename string) (string, string) {
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

// JoinNameExt returns the filename as name.ext.
// If the ext is empty the name is used as the filename.
func JoinNameExt(name, ext string) string {
	if ext == "" {
		return name
	}
	return name + "." + ext
}

// CreateDir shortcut to create a directory in the given path.
func CreateDir(path string) error {
	return os.MkdirAll(path, 0764)
}

// CreateDirPanic shortcut to create a directory in the given path.
// if an error occurs it panics.
func CreateDirPanic(path string) {
	if err := os.MkdirAll(path, 0764); err != nil {
		panic("Cant create directory " + path + ": " + err.Error())
	}
}
