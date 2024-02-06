package telegram

type Telegram struct {
	ApiKey string
}

func (t *Telegram) SendMessage(message string) bool {
	return true
}
