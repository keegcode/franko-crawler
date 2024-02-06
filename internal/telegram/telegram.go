package telegram

import (
	"net/http"
	"net/url"
)

type Telegram struct {
	ApiKey    string
	ChannelId string
}

func (t *Telegram) SendMessage(message string) error {
	uri := "https://api.telegram.org/bot" + t.ApiKey + "/sendMessage"
	query := "?text=" + url.QueryEscape(message) + "&" + "chat_id=" + t.ChannelId

	_, err := http.Get(uri + query)
	if err != nil {
		return err
	}

	return nil
}
