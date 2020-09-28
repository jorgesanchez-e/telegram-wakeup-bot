package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/jorgesanchez-e/telegram-wakeup-bot/pkg/bot"
	"github.com/jorgesanchez-e/telegram-wakeup-bot/pkg/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := bot.New(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	stopBot(bot, cancel)
}

func stopBot(bot bot.Bot, cancel context.CancelFunc) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Info("stopping bot")

	bot.Stop()
	log.Info("exit")
}
