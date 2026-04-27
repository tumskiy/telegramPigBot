package main

import (
	"math/rand"
	"os"
	"os/signal"
	"piggifbot/command"
	"piggifbot/command/htopd"
	"piggifbot/pkg/gifcache"
	"piggifbot/pkg/server"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var htoyaCaptions = map[string]string{
	"agent.gif":    "Ты агент 69",
	"billy.gif":    "Ты Билли",
	"clown.gif":    "Ты клоун",
	"don.gif":      "Ты Дон",
	"gigachad.gif": "Ты Гигачед",
	"it.gif":       "Ты ойтешнек",
	"obeme.gif":    "Ты черножопый",
	"ohmy.gif":     "Ты ценитель",
	"petro.gif":    "Ты призрак Киева",
	"rat.gif":      "Ля ты крыса конешн",
	"ryglo.gif":    "Ты рыгло",
	"serb.gif":     "Ты взял Жепу",
	"skyrim.gif":   "Ты пробудившийся",
	"sunboy.gif":   "Ты Санбой",
	"zmyh.gif":     "Ты Жмых",
	"dolboeb.gif":  "Ты сказочный долбоёб",
	"gey.gif":      "Ты гей",
	"girkin.gif":   "Ты предрекал мобилизацию",
	"harold.gif":   "Ты скрываешь боль",
	"kapibara.gif": "Ты капибаржа",
	"pohuist.gif":  "Ты похуист",
	"shok.gif":     "Ты в шоке",
	"sina.gif":     "Ты Сина (Джон)",
	"sanders.gif":  "Ты полковник Сандерс",
	"vasya.gif":    "Ты Васян",
	"zoltan.gif":   "Ты словил легендарОЧКУ, ты Золтан!!!",
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	srv := server.NewServer()
	bot := srv.BotApi

	gifCache := gifcache.NewCache(bot, "./.gifcache")

	if err := gifCache.PreloadCache(); err != nil {
		logrus.Errorf("Failed to preload: %v", err)
	}

	if err := gifCache.LoadAllGifsWithCaptions("./command/hru", "hru", nil); err != nil {
		logrus.Warnf("Failed to load hru gifs: %v", err)
	}

	if err := gifCache.LoadAllGifsWithCaptions("./command/htoya", "htoya", htoyaCaptions); err != nil {
		logrus.Warnf("Failed to load htoya gifs: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120

	updates := bot.GetUpdatesChan(u)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}

			switch update.Message.Command() {
			case "hru":
				random := rand.Intn(1000)
				if random > 100 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyToMessageID = update.Message.MessageID

					gif, err := gifCache.GetRandomGif("hru")
					if err != nil {
						logrus.Errorf("Failed to get gif from cache: %v", err)
						bot.Send(command.PigText(update.Message.Chat.ID, update.Message.MessageID))
						continue
					}

					animationMsg := tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileID(gif.TgFileID))
					animationMsg.ReplyToMessageID = update.Message.MessageID
					_, err = bot.Send(animationMsg)
					if err != nil {
						logrus.Errorf("Failed to send animation: %v", err)
					}
					continue
				}

			case "htoya":
				random := rand.Intn(1000)
				if random > 10 {
					gif, err := gifCache.GetRandomGif("htoya")
					if err != nil {
						logrus.Errorf("Failed to get gif from cache: %v", err)
						continue
					}

					animationMsg := tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileID(gif.TgFileID))
					animationMsg.ReplyToMessageID = update.Message.MessageID
					animationMsg.Caption = gif.Caption 
					_, err = bot.Send(animationMsg)
					if err != nil {
						logrus.Errorf("Failed to send animation: %v", err)
					}
					continue
				}

				zoltanGif, err := gifCache.GetGifByName("zoltan.gif")
				if err == nil {
					animationMsg := tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileID(zoltanGif.TgFileID))
					animationMsg.ReplyToMessageID = update.Message.MessageID
					animationMsg.Caption = zoltanGif.Caption 
					_, err = bot.Send(animationMsg)
					if err != nil {
						logrus.Errorf("Failed to send zoltan: %v", err)
					}
				}

			case "htopidoreg":
				u := &htopd.User{
					Name: update.Message.From.UserName,
					ID:   update.Message.From.ID,
				}
				_, err := bot.Send(u.UserReg(u.ID, update.Message.Chat.ID, update.Message.MessageID))
				if err != nil {
					logrus.Errorf("Error: %v", err)
				}

			case "htopidor":
				pd := &htopd.PD{
					UserID: update.Message.From.ID,
					ChatID: update.Message.Chat.ID,
				}
				_, err := bot.Send(pd.SendTodayPD(pd.UserID, pd.ChatID, update.Message.MessageID))
				if err != nil {
					logrus.Errorf("Error: %v", err)
				}

			case "skokapidorov":
				pd := &htopd.PD{
					UserID: update.Message.From.ID,
					ChatID: update.Message.Chat.ID,
				}
				_, err := bot.Send(pd.SendGetCountAllPD(pd.ChatID, update.Message.MessageID))
				if err != nil {
					logrus.Errorf("Error: %v", err)
				}
			}
		}
	}()

	<-sigChan
	logrus.Info("Shutting down...")
}
