package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"os"
)

const TOKEN = "5748604300:AAGluxwv-mlwv9Vcw88jPokzs_QIa1VnXYQ"
const (
	PIG = `
	⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠈⢿⣿⣿⣿⣿⣿⠋⠀⠀
	⢹⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣿⠏⠀⠀⠀⠀
	⠈⣿⡇⠀⠀⠀⡀⠀⠀⠀⠀⠀⠿⣿⣿⠀⠀⠀⠀⠀
	⡀⢹⣧⠀⠀⠰⠋⠁⠀⠀⠀⠀⠀⠈⣉⡐⠀⠀⠀⠀
	⠁⠈⣿⣧⠀⠀⠀👁⠀👁 ⢸⣿⣿⣿⠀⠀⠀⠀
	⡂⠀⢹⣿⣆⠀⠀⠀⠀⠀⠀   ⢠⣿⣿⣿⣿⣦⠀⠀⠀
	⠁⠀⠀⠻⣿⡆⠸⣶⣤⡀🐽⣸⣿⣿⣿⣿⣿⣷⣧⠀
	⠬⠀⠀⠀⢿⣷⠀⣻⣿⣿⣿🫧🫧⣿⣿⣿⣿⣿⣿⣧
	⣷⠁⢐⠀⠈⣿⣿⣿⣿⣿⣿⣿⣿🫧⣿⣿⣿⣿⣿⣿
	⣿⣶⣶⣦⠀⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`
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

			random := rand.Intn(1000)
			if random < 980 {
				// Отправляем гиф
				_, err = bot.Send(randomGifs(update.Message.Chat.ID))
				if err != nil {
					return
				}
				continue
			}
			// Отправляем текст
			_, err = bot.Send(pigText(update.Message.Chat.ID))
			if err != nil {
				return
			}
		}
	}
}

// Функция отпрвки гифки
func randomGifs(chatId int64) tgbotapi.DocumentConfig {
	//Массив с именами гифок из файла
	var gifs = []string{"1.gif", "2.gif", "3.gif"}
	//Рандомим по длине массива, индекс становится = имени гифки
	random := rand.Intn(len(gifs))
	//Форматтируем из  массива в стрингу-путь
	format := fmt.Sprintf("./gifs/%s", gifs[random])

	//Открываем файл по пути и возвращаем
	reader, _ := os.Open(format)
	file := tgbotapi.FileReader{
		Name:   format,
		Reader: reader,
	}
	return tgbotapi.NewDocument(chatId, file)
}

// Функция отправки стринговой свиньи
func pigText(chatId int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatId, PIG)
}
