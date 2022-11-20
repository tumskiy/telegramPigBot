package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"piggifbot/command"
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
		chatId := update.Message.Chat.ID
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			// #КОНЕЦ СТАНДАРТНОЙ БИБЛИОТЕКИ
			if update.Message.Command() == "hru" {

				random := rand.Intn(1000)
				if random > 100 {
					// Отправляем гиф
					_, err := bot.Send(command.RandomGifs(chatId))
					if err != nil {
						return
					}
					continue
				}
				// Отправляем текст
				_, err := bot.Send(command.PigText(chatId))
				if err != nil {
					return
				}
			}

			if update.Message.Command() == "htoya" {
				random := rand.Intn(1000)
				if random > 10 {
					_, err := bot.Send(command.HtoyaGifs(chatId))
					if err != nil {
						return
					}
					continue
				}
				_, err := bot.Send(command.SendZoltan(chatId))
				if err != nil {
					return
				}

			}
		}
	}
}
