package bot

import (
	"context"
	"fmt"

	"github.com/apex/log"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Config depics config interface
type Config interface {
	Webhook() string
	Cert() string
	Key() string
	Token() string
}

type Bot struct {
	ctx            context.Context
	webhook        string
	token          string
	cert           string
	key            string
	api            *botapi.BotAPI
	incomeMessages botapi.UpdatesChannel
}

// New returns new bot instance
func New(ctx context.Context, config Config) (Bot, error) {
	newBot := Bot{
		ctx:     ctx,
		token:   config.Token(),
		webhook: config.Webhook(),
		cert:    config.Cert(),
		key:     config.Key(),
	}

	if err := newBot.start(); err != nil {
		return Bot{}, err
	}

	return newBot, nil
}

func (b *Bot) start() error {
	var err error

	if b.api, err = botapi.NewBotAPI(b.token); err != nil {
		return err
	}

	b.api.Debug = true
	u := botapi.NewUpdate(0)
	u.Timeout = 60

	if b.incomeMessages, err = b.api.GetUpdatesChan(u); err != nil {
		return err
	}

	go b.driveMessages()

	return nil
}

func (b Bot) driveMessages() {
	for {
		select {
		case <-b.ctx.Done():
			return
		case msg := <-b.incomeMessages:
			log.WithFields(log.Fields{
				"msg": fmt.Sprintf("%#v", msg),
			}).Info("message")
		}
	}
}

// Stop stops bot
func (b Bot) Stop() {
	b.api.StopReceivingUpdates()
}
