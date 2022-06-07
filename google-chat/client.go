package google_chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

type Client interface {
	SendMessage(*Message) error
}

type client struct {
	url string
}

func NewClient(url string, key string, token string, conversationKey string) Client {
	fullURL := ""
	if conversationKey == "" {
		fullURL = url + "/messages?key=" + key + "&token=" + token
	} else {
		fullURL = url + "/messages?key=" + key + "&token=" + token + "&threadKey=" + conversationKey
	}
	return &client{
		url: fullURL,
	}
}

func (c *client) SendMessage(msg *Message) error {

	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	resp, err := http.Post(c.url, "application/json", buf)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("unable to post message %d", resp.StatusCode)
	}

	return nil
}
