package thirdparties

import (
	"bytes"
	"fmt"
	"net/http"
	configs "pink/settings/configs"
)

// Slack struct
type Slack struct{}

// PushNotification func
// Arguments:
// 1. message -> string
func (s Slack) PushNotification(message string) {
	config := configs.SlackConfig()
	webHook := config.WebHook
	mes := []byte(fmt.Sprintf(`{"text":"%s"}`, message))
	req, err := http.NewRequest("POST", webHook, bytes.NewBuffer(mes))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
