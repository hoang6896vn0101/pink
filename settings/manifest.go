/*
	Manifest
	- Author: Hoang Pham Dang
	- Content: Export path in system.
*/

package manifest

var envPath = "settings/env"
var mailerPath = "domains/mailers/"

// EnvConfigPath func
func EnvConfigPath(env string) string {
	return envPath + "/" + env + ".yaml"
}

// MailerConfigPath func
func MailerConfigPath(name string) string {
	return mailerPath + "/" + name + "html"
}
