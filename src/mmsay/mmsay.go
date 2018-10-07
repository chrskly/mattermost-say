
package mmsay

import (
    "os"
    "fmt"
    //"net/http"
    //"net/url"
    "mmwebhook"
)

func Say(message string) (string, error) {
    var c conf
    _, err := c.LoadConfig()
    if err != nil {
        return "", fmt.Errorf("problem loading config: %v", err)
    }
    fmt.Fprintf(os.Stdout, "configuration %v\n", c)
    fmt.Fprintf(os.Stdout, "Posting message to channel %v as user %v\n", c.DefaultChannel, c.User)

//    resp, err := http.PostForm(c.WebhookUrl,
//	url.Values{"username": {c.User}, "channel": {c.DefaultChannel}, "text": {message}})

    resp, err := mmwebhook.Post(c.User, c.Password, message, c.DefaultChannel, c.WebhookUrl)
    if err != nil {
        return "", fmt.Errorf("problem posting message: %v", err)
    }
    _ = resp
    return "", nil

}
