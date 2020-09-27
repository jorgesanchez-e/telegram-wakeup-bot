package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jorgesanchez-e/telegram-wakeup-bot/pkg/config"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", conf)
}
