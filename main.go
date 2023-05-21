package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nats-io/nats.go"

	"github.com/agrrh/tg-sender/internal/handler"
)

func main() {
	// Init

	var appName, tgToken, natsAddr, natsPrefix string
	var err error
	var ok bool

	appName, ok = os.LookupEnv("APP_NAME")
	if !ok {
		log.Panic("set application name")
	}

	tgToken, ok = os.LookupEnv("APP_TG_TOKEN")
	if !ok {
		log.Panic("can't start without Telegram token")
	}

	natsAddr, ok = os.LookupEnv("APP_NATS_ADDR")
	if !ok {
		natsAddr = nats.DefaultURL
	}

	natsPrefix, ok = os.LookupEnv("APP_NATS_PREFIX")
	if !ok {
		natsPrefix = "dummy"
	}

	// NATS

	nc, err := nats.Connect(natsAddr)
	if err != nil {
		log.Panic("could not connect: ", err)
	}

	defer nc.Close()
	defer nc.Flush()

	js, err := nc.JetStream()
	if err != nil {
		log.Panic("could not get jetstream: ", err)
	}

	natsStreamChannel := fmt.Sprintf("%s.tg.out", natsPrefix)

	js.AddStream(&nats.StreamConfig{
		Name: appName,
		Subjects: []string{
			natsStreamChannel,
			fmt.Sprintf("%s.*", natsStreamChannel),
		},
		Discard: nats.DiscardOld,
		MaxMsgs: 1000,
	})

	js.AddConsumer("worker", &nats.ConsumerConfig{})

	// Telegram

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Panic("could not get telegram: ", err)
	}

	log.Printf("authorized on account %s", bot.Self.UserName)

	// Handler

	h := handler.NewHandler(*bot)

	// Run logic

	js.Subscribe(
		fmt.Sprintf("%s.>", natsStreamChannel),
		h.Handle,
		nats.Durable("worker"),
	)

	for {
		time.Sleep(1 * time.Second)
	}
}
