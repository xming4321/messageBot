package helper

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"messageBot/conf"
)

var tgUpdateChan tgbotapi.UpdatesChannel
var tgBot *tgbotapi.BotAPI

func InitTgBot() {
	tgBotApi, err := tgbotapi.NewBotAPI(conf.Conf.TgBot.Token)
	if err != nil {
		PanicfLogger(nil, "Telegram api connect with token:%s error: %v", conf.Conf.TgBot.Token, err)
	}

	tgBotApi.Debug = conf.Conf.TgBot.Debug
	tgBot = tgBotApi

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	tgUpdateChan = tgBot.GetUpdatesChan(updateConfig)
}

func DealMessage(handle func(bot *tgbotapi.BotAPI, update tgbotapi.Update)) {
	for update := range tgUpdateChan {
		go handle(tgBot, update)
	}
}
