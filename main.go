package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
)

func main() {
	// #НАЧАЛО СТАНДАРТНОЙ БИБЛИОТЕКИ
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			// #КОНЕЦ СТАНДАРТНОЙ БИБЛИОТЕКИ

			// 25% на появление гифки или текста
			random := rand.Intn(100)
			if random < 75 {
				// Отправляем гиф
				bot.Send(randomGifs(update.Message.Chat.ID))
				continue
			}
			// Отправляем текст
			bot.Send(pigText(update.Message.Chat.ID))
		}
	}
}
