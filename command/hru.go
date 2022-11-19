package command

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"os"
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
func RandomGifs(chatId int64) tgbotapi.DocumentConfig {
	//Массив с именами гифок из файла
	var gifs = []string{"1.gif", "2.gif", "3.gif", "4.gif", "5.gif", "6.gif", "7.gif"}
	//Рандомим по длине массива, индекс становится = имени гифки
	random := rand.Intn(len(gifs))
	//Форматтируем из  массива в стрингу-путь
	format := fmt.Sprintf("./command/hru/%s", gifs[random])
	formatName := fmt.Sprintf("%s", gifs[random])

	//Открываем файл по пути и возвращаем
	reader, _ := os.Open(format)
	file := tgbotapi.FileReader{
		Name:   formatName,
		Reader: reader,
	}
	return tgbotapi.NewDocument(chatId, file)
}

// PigText Функция отправки стринговой свиньи
func PigText(chatId int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatId, "ТЕБЕ ПОВЕЗЛО!!!! ВОТ ТЕБЕ РЕДКИЙ ХРЯК\n"+PIG)
}