package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"os"
)

type FileConfig struct {
	name string
	msg  string
}

func HtoyaGifs(chatId int64) tgbotapi.DocumentConfig {
	workDir := "./command/htoya/"
	var massGifsCaptions = []FileConfig{
		{
			name: "agent.gif",
			msg:  "Ты агент 69",
		},
		{
			name: "billy.gif",
			msg:  "Ты Билли",
		},
		{
			name: "clown.gif",
			msg:  "Ты клоун",
		},
		{
			name: "don.gif",
			msg:  "Ты Дон",
		},
		{
			name: "gigachad.gif",
			msg:  "Ты Гигачед",
		},
		{
			name: "it.gif",
			msg:  "Ты ойтешнек",
		},
		{
			name: "obeme.gif",
			msg:  "Ты Абэме",
		},
		{
			name: "ohmy.gif",
			msg:  "Ты ценитель",
		},
		{
			name: "petro.gif",
			msg:  "Ты призрак Киева",
		},
		{
			name: "rat.gif",
			msg:  "Ля ты крыса конешн",
		},
		{
			name: "ryglo.gif",
			msg:  "Ты рыгло",
		},
		{
			name: "serb.gif",
			msg:  "Ты взял Жепу",
		},
		{
			name: "skyrim.gif",
			msg:  "Ты пробудившийся",
		},
		{
			name: "sunboy.gif",
			msg:  "Ты Санбой",
		},
		{
			name: "zmyh.gif",
			msg:  "Ты Жмых",
		},
	}
	//делаем массив строк равным длине массива объектов
	var namesArray = make([]string, len(massGifsCaptions))
	//Заполняем новый массив именем гифки
	for index, value := range massGifsCaptions {
		namesArray[index] = value.name
	}
	// Создаем файл
	name := RandomName(namesArray)
	reader, _ := os.Open(workDir + name)
	file := tgbotapi.FileReader{
		Name:   name,
		Reader: reader,
	}
	//Ищем описание по имени файла
	var captions string
	for _, massGifCaption := range massGifsCaptions {

		if name == massGifCaption.name {
			captions = massGifCaption.msg
			break
		}
	}
	//Выводим описание
	fileConfig := tgbotapi.NewDocument(chatId, file)
	fileConfig.Caption = captions
	return fileConfig
}

// RandomName Рандомайзер массива
func RandomName(str []string) string {
	random := rand.Intn(len(str))
	name := str[random]
	return name
}

func SendZoltan(chatId int64) tgbotapi.DocumentConfig {
	reader, _ := os.Open("./command/htoya/zoltan.gif")
	file := tgbotapi.FileReader{
		Name:   "zoltan.gif",
		Reader: reader,
	}
	fileConfig := tgbotapi.NewDocument(chatId, file)
	fileConfig.Caption = "Ты словил легендарОЧКУ, ты Золтан!!!"
	return fileConfig
}
