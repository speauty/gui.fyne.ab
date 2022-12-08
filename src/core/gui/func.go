package gui

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetRuntimeDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}
