package slack

import (
	"bytes"
	"fmt"
	"net/http"
	"pink/config/env"
)

// PushNotification func
// Arguments:
// 1. message -> string
func PushNotification(message string) {
	config := env.SlackConfig()
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
