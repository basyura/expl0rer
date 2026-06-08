package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestToolKey(t *testing.T) {
	tests := []struct {
		name string
		exe  string
		want string
	}{
		{
			name: "windows absolute path",
			exe:  `C:\Program Files\vim\gvim.exe`,
			want: "gvim",
		},
		{
			name: "plain executable name",
			exe:  "explorer.exe",
			want: "explorer",
		},
		{
			name: "path without extension",
			exe:  filepath.Join("tools", "custom"),
			want: "custom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToolKey(tt.exe)
			if got != tt.want {
				t.Fatalf("ToolKey(%q) = %q, want %q", tt.exe, got, tt.want)
			}
		})
	}
}

func TestResolveToolExe(t *testing.T) {
	cfg := Config{
		Tools: map[string]string{
			"gvim": `D:\tools\gvim.exe`,
		},
	}

	got := resolveToolExe(`C:\Program Files\vim\gvim.exe`, cfg)
	want := `D:\tools\gvim.exe`
	if got != want {
		t.Fatalf("resolveToolExe() = %q, want %q", got, want)
	}
}

func TestResolveToolExeReturnsDefaultWhenMissing(t *testing.T) {
	cfg := Config{
		Tools: map[string]string{
			"gvim": `D:\tools\gvim.exe`,
		},
	}

	want := "explorer.exe"
	got := resolveToolExe(want, cfg)
	if got != want {
		t.Fatalf("resolveToolExe() = %q, want %q", got, want)
	}
}

func TestLoadFile(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "expl0rer.toml")
	content := "[tools]\n" +
		"gvim = 'C:\\Program Files\\vim\\gvim.exe'\n"
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := loadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}

	got := cfg.Tools["gvim"]
	want := `C:\Program Files\vim\gvim.exe`
	if got != want {
		t.Fatalf("cfg.Tools[%q] = %q, want %q", "gvim", got, want)
	}
}
