// server.go
package server

import (
	"piggifbot/env"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var token = env.ParseEnv("TOKEN")

type Server struct {
	BotApi     *tgbotapi.BotAPI
	HttpClient *HttpClientBot
}

func NewServer() *Server {
	httpClient := GetHTTPClient()
	bot, err := tgbotapi.NewBotAPIWithClient(
		token,
		tgbotapi.APIEndpoint,
		httpClient.Client,
	)
	if err != nil {
		logrus.Panicf("Failed to init Telegram Bot: %v", err)
	}

	bot.Debug = true
	logrus.Infof("Бот авторизован: %s", bot.Self.UserName)
	logrus.Info(GetProxyInfo())

	return &Server{
		BotApi:     bot,
		HttpClient: httpClient,
	}
}
