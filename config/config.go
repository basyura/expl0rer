package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Tools map[string]string `toml:"tools"`
}

func ResolveToolExe(exe string) string {
	cfg, err := Load()
	if err != nil {
		return exe
	}
	return resolveToolExe(exe, cfg)
}

func resolveToolExe(exe string, cfg Config) string {
	if toolExe := cfg.Tools[ToolKey(exe)]; toolExe != "" {
		return toolExe
	}
	return exe
}

func Load() (Config, error) {
	exePath, err := os.Executable()
	if err != nil {
		return Config{}, err
	}

	configPath := filepath.Join(filepath.Dir(exePath), "expl0rer.toml")
	return loadFile(configPath)
}

func loadFile(configPath string) (Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func ToolKey(exe string) string {
	base := filepath.Base(exe)
	return strings.TrimSuffix(base, filepath.Ext(base))
}
