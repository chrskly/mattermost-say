
package mmwebhook

func Post(username string, message string, channel string, webhookUrl string) error {
    // Post a message to a channel using a webhook
    resp, err := http.PostForm(webhookUrl,
                   url.Values{"username": username, "channel": channel, "text": message})
}
