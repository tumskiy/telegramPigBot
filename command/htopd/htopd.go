package htopd

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PDSender interface {
	UserReg()
	SendTodayPD(int64, int64, int) tgbotapi.MessageConfig
}

type Databaser interface {
	CrateUser(int64) (bool, error)
	GetUserName(int64) (*User, error)

	CreateTodayPD() error
	GetCountAllPD() (int64, int, error)
}

type User struct {
	ID   int64
	Name string
}

type PD struct {
	ID     int64
	UserID int64
	ChatID int64
	Date   time.Time
}

type UserCount struct {
	Name  string
	Count int
}

func (pd *PD) SendTodayPD(userID, chatID int64, replay int) tgbotapi.MessageConfig {
	create, victim, err := pd.CreateTodayPD()
	if err != nil {
		send := sendError(chatID, replay, err)
		return send
	}

	user := &User{}
	user, err = user.GetUserByID(victim)
	if err != nil {
		send := sendError(chatID, replay, err)
		return send
	}

	if create {
		messageConfig := tgbotapi.NewMessage(chatID, fmt.Sprintf("Опа, пидор дня сегодня: @%s", user.Name))
		messageConfig.ReplyToMessageID = replay
		return messageConfig
	} else {
		messageConfig := tgbotapi.NewMessage(chatID, fmt.Sprintf("Сегодняшний пидор уже был найден, это: @%s", user.Name))
		messageConfig.ReplyToMessageID = replay
		return messageConfig
	}
}

func (pd *PD) SendGetCountAllPD(chatID int64, reply int) tgbotapi.MessageConfig {
	u := &User{}
	users, err := u.GetAllUsers()
	if err != nil {
		send := sendError(chatID, reply, err)
		return send
	}

	userCounts := make(map[string]int)

	for _, user := range users {
		userID := user.ID
		count, err := pd.GetCountPDByUserID(userID)
		if err != nil {
			send := sendError(chatID, reply, err)
			return send
		}

		userCounts[user.Name] = count
	}

	// Преобразование map в slice для сортировки
	type kv struct {
		Key   string
		Value int
	}
	var sortedUserCounts []kv
	for k, v := range userCounts {
		sortedUserCounts = append(sortedUserCounts, kv{k, v})
	}
	sort.Slice(sortedUserCounts, func(i, j int) bool {
		return sortedUserCounts[i].Value > sortedUserCounts[j].Value
	})

	var preMsg []string
	for _, userCount := range sortedUserCounts {
		msg := fmt.Sprintf("%s - %d раз\n", userCount.Key, userCount.Value)
		preMsg = append(preMsg, msg)
	}

	messageText := strings.Join(preMsg, "")

	messageConfig := tgbotapi.NewMessage(chatID, fmt.Sprintf("Список хээв:\n%s", messageText))
	messageConfig.ReplyToMessageID = reply
	return messageConfig
}

func (u *User) UserReg(userID, chatID int64, replay int) tgbotapi.MessageConfig {

	err := u.CreateUser()
	if err != nil {
		send := sendError(chatID, replay, err)
		return send
	}
	_, err = u.GetUserByID(userID)
	if err != nil {
		send := sendError(chatID, replay, err)
		return send
	}
	messageConfig := tgbotapi.NewMessage(chatID, fmt.Sprintf("Теперь Вы в списке будущих пидоров, @%s", u.Name))
	messageConfig.ReplyToMessageID = replay
	return messageConfig
}

func sendError(chatID int64, replay int, err error) tgbotapi.MessageConfig {
	messageConfig := tgbotapi.NewMessage(chatID, fmt.Sprintf("Чел, у нас ошибка: [%s]", err))
	messageConfig.ReplyToMessageID = replay
	return messageConfig
}
