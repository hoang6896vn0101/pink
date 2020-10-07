package configs

// Mailer struct
type Mailer struct {
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Domain   string `yaml:"domain"`
	Port     string `yaml:"port"`
}

// MailerConfig func
func MailerConfig() Mailer {
	return getConfig("development").Mailer
}
