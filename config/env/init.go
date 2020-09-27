package env

import (
	filepath "path/filepath"
	infra "pink/infrastructure"
	lib "pink/libs"
)

// Config struct
type Config struct {
	Mysql struct {
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		UserName     string `yaml:"user_name"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"database_name"`
	}
}

var config *Config

// GetConfig func
// Arguments:
// - No
// Return:
// 1. MySQLConfig struct
func GetConfig(env string) *Config {
	if config == nil {
		var temp Config
		path, err := filepath.Abs(configPath(env))
		if err != nil {
			panic("Error path file")
		}
		lib.YAMLParse(&temp, path)
		config = &temp
	}
	return config
}

// configPath
// Arguments:
// 1. env => Env of instance
// Return:
// 1. path => path of config
func configPath(env string) string {
	return infra.ConfigPath + env + ".yaml"
}
