package handler

import (
	"log"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nats-io/nats.go"

	"github.com/agrrh/tg-sender/internal/types"
)

type Handler struct {
	TgBot tgbotapi.BotAPI
}

func NewHandler(tgBot tgbotapi.BotAPI) *Handler {
  return &Handler{TgBot: tgBot}
}

func (h Handler) Handle(m *nats.Msg) {
	var err error

	r := types.Reply{}

	err = json.Unmarshal(m.Data, &r)
	if err != nil {
		log.Printf("could not decode message data %s", string(m.Data))
	}

	mReply := tgbotapi.NewMessage(r.Chat, r.Text)
	mReply.ReplyToMessageID = r.ReplyTo

	_, err = h.TgBot.Send(mReply)
	if err != nil {
		log.Printf("could not send message %s", mReply)
	}
}
