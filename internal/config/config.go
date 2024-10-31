package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

var MapLog = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

var Configs Config

type Config struct {
	Server server   `yaml:"server"`
	Logger logger   `yaml:"logger"`
	DB     database `yaml:"database"`
}

type server struct {
	Port uint `yaml:"port"`
}

type logger struct {
	Level string `yaml:"level"`
}

type database struct {
	DSN string `yaml:"dsn"`
}

func Load(path string) error {
	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(file, &Configs)
}
