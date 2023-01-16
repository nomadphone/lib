package users

type User struct {
	TelegramUsername string `json:"telegramUsername"`
	Name             string `json:"name"`
	TwillioPhone     string `json:"twillioPhone"`
	TelegramChatID   int64  `json:"telegramChatID"`
}
