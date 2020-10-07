package env

// Mysql Config
type Mysql struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	UserName     string `yaml:"user_name"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"database_name"`
}

// MySQLConfig func
func MySQLConfig() Mysql {
	return getConfig("development").Mysql
}
