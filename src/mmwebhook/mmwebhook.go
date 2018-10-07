
package mmwebhook

import (
    "net/http"
    //"net/url"
    "fmt"
    "os"
    "bytes"
)

func Post(username *string, password *string, message string, channel *string, webhookUrl *string) (string, error) {
    // Post a message to a channel using a webhook
    //resp, err := http.PostForm(*webhookUrl,
    //    url.Values{"username": {*username}, "channel": {*channel},
    //               "text": {message}})
    client := &http.Client{}

    // Set up POST data
    payload_str := fmt.Sprintf("{\"username\":\"%s\",\"channel\":\"%s\",\"text\":\"%s\"}", *username, *channel, message)
    fmt.Fprintf(os.Stdout, "payload : %s\n", payload_str)
    var payload = []byte(payload_str)

    req, err := http.NewRequest("POST", *webhookUrl, bytes.NewBuffer(payload))
    if err != nil {
        return "", err
    }

    // Set up authentication
    req.SetBasicAuth(*username, *password)

    // Execute request
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    fmt.Fprintf(os.Stdout, "resp %v", resp)
    return "", nil
}
