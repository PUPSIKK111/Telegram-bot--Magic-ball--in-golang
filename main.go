package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Token = "Your TOKEN/Твой токен"

var bot *tgbotapi.BotAPI

var allNames = [2]string{"брайн", "смарт"}

var chatId int64

var answers = []string{
	"Конечно же нет",
	"Без сомнений!",
	"Конечно же да",
	"Неуверен, но скорее всего да!",
	"Неуверен, но скорее всего нет!",
	"Думаю что нет, но чуйка говорит что да!",
	"Если задонатишь мне сто рублей, то да!",
	"Если не задонатишь мне сто рублей, то нет!",
	"Если задонатишь мне сто рублей, то нет!",
	"Если не задонатишь мне сто рублей, то да!",
	"Скоре да, чем нет!",
	"Скоре нет, чем да!",
	"100% нет, без вариантов",
	"100% да, без вариантов",
}

func getFortunTellersAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func ConnectionWithTg() {
	var err error
	bot, err = tgbotapi.NewBotAPI(Token)
	if err != nil {
		panic("Cannot connect to telegram bot hehehehhehehee")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}

func isMessageForMyBot(update *tgbotapi.Update) bool {
	if update.Message == nil && update.Message.Text == "" {
		return false
	}
	msgTextCase := strings.ToLower(update.Message.Text)
	for _, name := range allNames {
		if strings.Contains(msgTextCase, name) {
			return true
		}

	}
	return false
}

func sendAnsew(update *tgbotapi.Update) { //пятсядва
	msg := tgbotapi.NewMessage(chatId, getFortunTellersAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	ConnectionWithTg()
	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID
			sendMessage("Введи текст который ты хочешь что бы вывелся при написании слова /start|Enter the text you want to appear when someone write the word /start")
		}
		if isMessageForMyBot(&update) {
			sendAnsew(&update)
		}
	}
}
