package main

import (
	"fmt"
	"log"
	"math/rand"
	"piggifbot/command"
	"piggifbot/command/htopd"
	"piggifbot/env"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	parseEnv := env.ParseEnv("TOKEN")
	// #НАЧАЛО СТАНДАРТНОЙ БИБЛИОТЕКИ
	bot, err := tgbotapi.NewBotAPI(parseEnv)
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

			if update.Message.Command() == "htopidoreg" {
				u := &htopd.User{
					Name: update.Message.From.UserName,
					ID:   update.Message.From.ID,
				}
				_, err := bot.Send(u.UserReg(u.ID, update.Message.Chat.ID, update.Message.MessageID))
				if err != nil {
					fmt.Println("Error:[%s]", err)
				}

			}

			if update.Message.Command() == "htopidor" {
				pd := &htopd.PD{
					UserID: update.Message.From.ID,
					ChatID: update.Message.Chat.ID,
				}
				_, err := bot.Send(pd.SendTodayPD(pd.UserID, pd.ChatID, update.Message.MessageID))
				if err != nil {
					fmt.Println("Error:[%s]", err)
				}

			}

			if update.Message.Command() == "skokapidorov" {
				pd := &htopd.PD{
					UserID: update.Message.From.ID,
					ChatID: update.Message.Chat.ID,
				}
				_, err := bot.Send(pd.SendGetCountAllPD(pd.ChatID, update.Message.MessageID))
				if err != nil {
					fmt.Println("Error:[%s]", err)
				}

			}
		}
	}
}

//привет
