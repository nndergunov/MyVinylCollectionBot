package api

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *ChatBot) Handle() {
	b.bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		b.logMessage(update.Message)

		switch update.Message.Command() {
		case "start":
			go b.startHandler(update)
		case "help":
			go b.helpHandler(update)
		default:
			go b.otherHandler(update)
		}
	}
}

func (b ChatBot) startHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("%v, the bot had started successfully", update.Message.Chat.FirstName)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	msg.ReplyMarkup = b.startButtons

	if _, err := b.bot.Send(msg); err != nil {
		b.log.ErrorLog(err)
	}
}

func (b ChatBot) helpHandler(update tgbotapi.Update) {
	text := "help" // todo help handler.

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		b.log.ErrorLog(err)
	}
}

func (b ChatBot) otherHandler(update tgbotapi.Update) {
	switch {
	case update.Message.Voice != nil:
		b.voiceHandler(update)
	case update.Message.VideoNote != nil:
		b.videoNoteHandler(update)
	case update.Message.Photo != nil:
		b.photoHandler(update)
	case update.Message.Video != nil:
		b.videoHandler(update)
	case update.Message.Poll != nil:
		b.pollHandler(update)
	case update.Message.Text != "":
		b.textHandler(update)
	}
}

func (b ChatBot) textHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've recieved \"%s\"", update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}