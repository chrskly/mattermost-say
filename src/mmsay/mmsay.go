
package mmsay

import (
    "fmt"
    "net/http"
    "net/url"
)

func Say(message string) (string, error) {
    var c conf
    _, err := c.LoadConfig()
    if err != nil {
        return "", fmt.Errorf("problem loading config: %v", err)
    }
    fmt.Println("configuration %v", c)
    fmt.Println("Posting message to channel %s as user %s", c.User, c.DefaultChannel)
    resp, err := http.PostForm(c.WebhookUrl,
	url.Values{"username": {c.User}, "channel": {c.DefaultChannel}, "text": {message}})
    if err != nil {
        return "", fmt.Errorf("problem posting message: %v", err)
    }
    _ = resp
    return "", nil

}
