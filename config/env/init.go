package env

import (
	filepath "path/filepath"
	infra "pink/infrastructure"
	lib "pink/libs"
	"sync"
)

// Mysql Config
type Mysql struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	UserName     string `yaml:"user_name"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"database_name"`
}

// Slack struct
type Slack struct {
	WebHook string `yaml:"web_hook"`
}

// Mailer struct
type Mailer struct {
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

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
func GetConfig(env string) *Config {
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

// MySQLConfig func
func MySQLConfig() Mysql {
	return GetConfig("development").Mysql
}

// SlackConfig func
func SlackConfig() Slack {
	return GetConfig("development").Slack
}

// MailerConfig func
func MailerConfig() Mailer {
	return GetConfig("development").Mailer
}
