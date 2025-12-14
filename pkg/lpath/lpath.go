package lpath

import (
	"path/filepath"
	"runtime"
	"strings"
)

var (
    ROOT_PATH string = getRootPath()
)

func getRootPath() string {
    var currentFileRelativePath = "/pkg/lpath/lpath.go"

    _, file, _, _ := runtime.Caller(0)
    index := strings.LastIndex(file, currentFileRelativePath)

    return file[:index]
}

func GetAbsolutetPath(relativePath string) string {
    return Join(ROOT_PATH, relativePath)
}

func Join(elem ...string) string {
    return filepath.Join(elem...)
}
