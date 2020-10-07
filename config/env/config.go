package env

import (
	filepath "path/filepath"
	infra "pink/infrastructure"
	lib "pink/libs"
	"sync"
)

// Config struct
type Config struct {
	Mysql
	Slack
	Mailer
}

var config *Config
var mu sync.Mutex

// GetConfig func
// Arguments:
// - No
// Return:
// 1. MySQLConfig struct
func getConfig(env string) *Config {
	mu.Lock()
	defer mu.Unlock()
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
