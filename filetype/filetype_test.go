package filetype

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsDirectory(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "file.txt")
	if err := os.WriteFile(filePath, []byte("text"), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "empty path",
			path: "",
			want: true,
		},
		{
			name: "directory",
			path: tempDir,
			want: true,
		},
		{
			name: "file",
			path: filePath,
			want: false,
		},
		{
			name: "missing path",
			path: filepath.Join(tempDir, "missing"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDirectory(tt.path)
			if got != tt.want {
				t.Fatalf("IsDirectory(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func TestIsTextFile(t *testing.T) {
	tempDir := t.TempDir()

	hiddenPath := filepath.Join(tempDir, ".env")
	writeFile(t, hiddenPath, []byte{0})

	noExtPath := filepath.Join(tempDir, "README")
	writeFile(t, noExtPath, []byte{0})

	knownExtPath := filepath.Join(tempDir, "notes.md")
	writeFile(t, knownExtPath, []byte{0})

	utf8Path := filepath.Join(tempDir, "unknown.ext")
	writeFile(t, utf8Path, []byte("hello\n"))

	binaryPath := filepath.Join(tempDir, "image.bin")
	writeFile(t, binaryPath, []byte{0x89, 'P', 'N', 'G', 0})

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "empty path",
			path: "",
			want: false,
		},
		{
			name: "hidden file",
			path: hiddenPath,
			want: true,
		},
		{
			name: "file without extension",
			path: noExtPath,
			want: true,
		},
		{
			name: "known text extension",
			path: knownExtPath,
			want: true,
		},
		{
			name: "utf8 content",
			path: utf8Path,
			want: true,
		},
		{
			name: "binary content",
			path: binaryPath,
			want: false,
		},
		{
			name: "missing path",
			path: filepath.Join(tempDir, "missing.bin"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsTextFile(tt.path)
			if got != tt.want {
				t.Fatalf("IsTextFile(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func writeFile(t *testing.T, path string, data []byte) {
	t.Helper()

	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
}