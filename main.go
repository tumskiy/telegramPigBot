package main

import (
	"log"
	"math/rand"
	"piggifbot/command"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "5874145224:AAGpGHS-_4LcyaQikUUi02UZ-_lQfIqSu2s"

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
			// #КОНЕЦ СТАНДАРТНОЙ БИБЛИОТЕКИ
			if update.Message.Command() == "hru" {
				random := rand.Intn(1000)
				if random > 100 {
					// Отправляем гиф
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyToMessageID = update.Message.MessageID
					_, err := bot.Send(command.RandomGifs(update.Message.Chat.ID, update.Message.MessageID))
					if err != nil {
						return
					}
					continue
				}
			}

			if update.Message.Command() == "htoya" {
				random := rand.Intn(1000)
				if random > 10 {
					_, err := bot.Send(command.HtoyaGifs(update.Message.Chat.ID, update.Message.MessageID))
					if err != nil {
						return
					}
					continue
				}
				_, err := bot.Send(command.SendZoltan(update.Message.Chat.ID, update.Message.MessageID))
				if err != nil {
					return
				}

			}
		}
	}
}

//привет
