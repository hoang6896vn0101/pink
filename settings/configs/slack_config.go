package configs

// Slack struct
type Slack struct {
	WebHook string `yaml:"web_hook"`
}

// SlackConfig func
func SlackConfig() Slack {
	return getConfig("development").Slack
}
