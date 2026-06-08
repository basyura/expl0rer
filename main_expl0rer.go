package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"expl0rer/config"
	"expl0rer/filetype"
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
	ext := strings.ToLower(filepath.Ext(path))
	if filetype.IsDirectory(path) {
		exe = `C:\Program Files\HmFilerClassic\HmFilerClassic.exe`
	} else if ext == ".sln" {

	} else if filetype.IsTextFile(path) {
		exe = `C:\Program Files\vim\gvim.exe`
	}

	exe = config.ResolveToolExe(exe)

	cmd := exec.Command(exe, path)
	if err := cmd.Run(); err != nil {
		if fmt.Sprintf("%s", err) != "exit status 1" {
			fmt.Println(err)
		}
	}
}
