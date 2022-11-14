package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"os"
)

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
		Name:   "2.gif",
		Reader: reader,
	}
	return tgbotapi.NewDocument(chatId, file)
}

// Функция отправки стринговой свиньи
func pigText(chatId int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatId, PIG)
}
