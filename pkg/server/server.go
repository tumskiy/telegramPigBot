package server

import (
	"piggifbot/env"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var token = env.ParseEnv("TOKEN")

type Server struct {
	botApi     *tgbotapi.BotAPI
	httpClient *HttpClientBot
}

func NewServer() *Server {
	httpClient := NewHTTPClient()
	bot, err := tgbotapi.NewBotAPIWithClient(
		token,
		tgbotapi.APIEndpoint,
		GetHTTPClient().Client,
	)
	if err != nil {
		logrus.Panicf("Failed to init Telegram Bot: %v", err)
	}

	bot.Debug = true
	logrus.Infof("Бот авторизован: %s", bot.Self.UserName)
	logrus.Info(GetProxyInfo())

	return &Server{
		botApi:     bot,
		httpClient: httpClient,
	}
}
