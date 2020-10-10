package configs

import (
	filepath "path/filepath"
	lib "pink/libs"
	manifest "pink/settings"
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
		path, err := filepath.Abs(manifest.EnvConfigPath(env))
		if err != nil {
			panic(err.Error())
		}
		lib.YAMLParse(&temp, path)
		config = &temp
	}
	return config
}
