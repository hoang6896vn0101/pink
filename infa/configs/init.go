package configs

import (
	"path/filepath"
	lib "pink/libs"
)

var (
	configPath = "infa/configs/"
)

// Config struct
type Config struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		UserName string `yaml:"user_name"`
		Password string `yaml:"password"`
	}
}

// GetConfig func
// Arguments:
// - No
// Return:
// 1. MySQLConfig struct
func GetConfig(env string) Config {
	var config Config
	configPath, err := filepath.Abs(pathConfig(env))
	if err != nil {
		panic("Error path file")
	}
	lib.YAMLParse(&config, configPath)
	return config
}

// pathConfig
// Arguments:
// 1. env => Env of instance
// Return:
// 1. path => path of config
func pathConfig(env string) string {
	return configPath + env + ".yaml"
}
