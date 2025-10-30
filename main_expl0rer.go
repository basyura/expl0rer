package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// 代表的なテキスト拡張子（小文字）
var textExts = map[string]struct{}{
	".txt": {}, ".md": {}, ".markdown": {},
	".go": {}, ".py": {}, ".rb": {}, ".php": {}, ".pl": {},
	".c": {}, ".h": {}, ".cpp": {}, ".cc": {}, ".hpp": {}, ".rs": {}, ".java": {}, ".cs": {}, ".kt": {},
	".js": {}, ".jsx": {}, ".ts": {}, ".tsx": {}, ".mjs": {}, ".cjs": {},
	".json": {}, ".yaml": {}, ".yml": {}, ".toml": {}, ".ini": {}, ".cfg": {}, ".conf": {}, ".properties": {},
	".xml": {}, ".html": {}, ".htm": {}, ".css": {}, ".scss": {}, ".less": {},
	".sh": {}, ".bash": {}, ".zsh": {}, ".bat": {}, ".ps1": {},
	".sql": {}, ".csv": {}, ".tsv": {}, ".log": {}, ".tex": {}, ".r": {},
	".gitignore": {}, ".gitattributes": {}, ".editorconfig": {},
}

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
	if isDirectory(path) {
		exe = `C:\Program Files\HmFilerClassic\HmFilerClassic.exe`
	} else if ext == ".sln" {

	} else if isTextFile(path) {
		exe = `C:\Program Files\vim\gvim.exe`
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

func isTextFile(path string) bool {
	if path == "" {
		return false
	}

	// 先頭が '.' の隠しファイルはテキスト扱い
	base := filepath.Base(path)
	if base != "." && base != ".." && strings.HasPrefix(base, ".") {
		return true
	}

	ext := strings.ToLower(filepath.Ext(path))
	// ユーザー要望: 拡張子なしはテキスト扱い
	if ext == "" {
		return true
	}

	// 代表的なテキスト拡張子の簡易判定（即 true）
	if _, ok := textExts[ext]; ok {
		return true
	}

	// 内容スニッフィング
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 8192)
	n, _ := f.Read(buf)
	if n == 0 {
		// 空ファイルはテキスト扱い
		return true
	}
	buf = buf[:n]

	// NUL バイトを含む場合はバイナリ判定
	if bytes.IndexByte(buf, 0) >= 0 {
		return false
	}

	// UTF-8 妥当ならテキスト
	if utf8.Valid(buf) {
		return true
	}

	// ASCII 可読文字の割合で判定
	printable := 0
	for _, b := range buf {
		if b == 9 || b == 10 || b == 13 || (b >= 32 && b < 127) {
			printable++
		}
	}
	ratio := float64(printable) / float64(len(buf))
	return ratio > 0.85
}
