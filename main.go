package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"piggifbot/hru"
)

const TOKEN = "5748604300:AAGluxwv-mlwv9Vcw88jPokzs_QIa1VnXYQ"

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

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "hru":
				random := rand.Intn(1000)
				if random < 900 {
					// Отправляем гиф
					_, err = bot.Send(hru.RandomGifs(update.Message.Chat.ID))
					if err != nil {
						return
					}
					continue
				}
				// Отправляем текст
				_, err = bot.Send(hru.PigText(update.Message.Chat.ID))
				if err != nil {
					return
				}
			case "htoya":
				tgbotapi.NewMessage(update.Message.Chat.ID, "Давай посмотрим, кто ты сейчас: `В разработке`")
			default:
				msg.Text = "Введи другую команду, я не знаю эту"
			}

		}
	}
}
