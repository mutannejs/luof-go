package lpath

import (
	"os/exec"
	"strings"
)

var (
    ROOT_PATH string = getRootPath()
)

func getRootPath() string {
    var cmdOut []byte
    var err error

    cmdOut, err = exec.Command("git", "rev-parse", "--show-toplevel").Output()

    if err == nil {
        return strings.TrimSpace(string(cmdOut)) + "/"
    }

    return ""
}

func GetAbsolutetPath(relativePath string) string {
    return ROOT_PATH + relativePath
}
