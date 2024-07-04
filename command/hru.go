package command

import (
	"fmt"
	"math/rand"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

// RandomGifs Функция отпрвки гифки
func RandomGifs(chatId int64, replay int) tgbotapi.DocumentConfig {

	var gifs = []string{"1.gif", "2.gif", "3.gif", "4.gif", "5.gif", "6.gif", "7.gif"}

	gifCache := make(map[int]string)
	for i, gif := range gifs {
		gifCache[i] = fmt.Sprintf("./command/hru/%s", gif)
	}

	random := rand.Intn(len(gifs))
	format := gifCache[random]
	formatName := gifs[random]

	reader, _ := os.Open(format)
	file := tgbotapi.FileReader{
		Name:   formatName,
		Reader: reader,
	}
	fileConfig := tgbotapi.NewDocument(chatId, file)
	fileConfig.ReplyToMessageID = replay
	return fileConfig
}

// PigText Функция отправки стринговой свиньи
func PigText(chatId int64, replay int) tgbotapi.MessageConfig {
	fileConfig := tgbotapi.NewMessage(chatId, "ТЕБЕ ПОВЕЗЛО!!!! ВОТ ТЕБЕ РЕДКИЙ ХРЯК\n"+PIG)
	fileConfig.ReplyToMessageID = replay
	return fileConfig
}
