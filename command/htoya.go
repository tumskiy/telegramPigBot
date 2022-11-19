package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
)

func RandomName(chatId int64) tgbotapi.MessageConfig {
	random := rand.Intn(5)
	if random == 1 {
		return tgbotapi.NewMessage(chatId, "1")
	}
	if random == 2 {
		return tgbotapi.NewMessage(chatId, "2")
	}
	if random == 3 {
		return tgbotapi.NewMessage(chatId, "3")
	}
	if random == 4 {
		return tgbotapi.NewMessage(chatId, "4")
	}
	return tgbotapi.NewMessage(chatId, "5")
}
