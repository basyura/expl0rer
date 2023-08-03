package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	path := ""
	if len(os.Args) >= 2 {
		path = os.Args[1]
		abs, err := filepath.Abs(path)
		if err != nil {
			fmt.Println(err)
		} else {
			path = abs
		}
	}

	exe := "explorer.exe"
	if isDirectory(path) {
		exe = `C:\Program Files\HmFilerClassic\HmFilerClassic.exe`
	}

	cmd := exec.Command(exe, path)
	if err := cmd.Run(); err != nil {
		if fmt.Sprintf("%s", err) != "exit status 1" {
			fmt.Println(err)
		}
	}
}

func isDirectory(path string) bool {

	if path == "" {
		return true
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
