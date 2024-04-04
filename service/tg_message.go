package service

import (
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"messageBot/helper"
	"messageBot/model"
	"time"
)

type TgMessageService struct{}

func NewTgMessageService(ctx *gin.Context) (*TgMessageService, error) {
	return new(TgMessageService), nil
}

/**
 * 创建
 */
func (*TgMessageService) Create(ctx *gin.Context, tgMessage model.TgMessage) (com *model.TgMessage, err error) {
	return model.InsertTgMessage(ctx, tgMessage)
}

/**
 * 获取
 */
func (*TgMessageService) GetById(ctx *gin.Context, id int) (com *model.TgMessage, err error) {
	return model.GetTgMessage(ctx, id)
}

/**
 * 删除
 */
func (*TgMessageService) Remove(ctx *gin.Context, id int) error {
	return nil
}

/**
 * 更新
 */
func (*TgMessageService) Update(ctx *gin.Context, TgMessage model.TgMessage) error {
	_, err := model.UpdateTgMessage(ctx, TgMessage)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 获取列表
 */
func (*TgMessageService) GetList(ctx *gin.Context, page int, pageSize int) (TgMessageList []model.TgMessage, err error) {
	offset := (page - 1) * pageSize
	return model.GetTgMessageList(ctx, pageSize, offset)
}

/**
 * 获取分页信息
 */
func (*TgMessageService) GetPageInfo(ctx *gin.Context, pageSize int) (totalCount int, totalPage int, err error) {
	totalCount, err = model.CountTgMessageList(ctx)
	totalPage = totalCount / pageSize
	if totalCount%pageSize != 0 {
		totalPage += 1
	}
	return totalCount, totalPage, err
}

func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	ctx, _ := gin.CreateTestContext(nil)
	nowTime := time.Now()
	chatID := update.Message.Chat.ID
	msg := model.TgMessage{
		ChatID:      chatID,
		UID:         update.Message.From.ID,
		IsCommand:   update.Message.IsCommand(),
		Command:     update.Message.Command(),
		TextContent: &update.Message.Text,
		SendTime:    time.Unix(int64(update.Message.Date), 0),
		ReceiveTime: nowTime,
		CreateTime:  nowTime,
		UpdateTime:  nowTime,
	}
	returnMsg, err := model.InsertTgMessage(ctx, msg)
	msg = *returnMsg
	if err != nil {
		helper.ErrorfLogger(ctx, "insert message error%+v", err)
		return
	}

	tempService, err := NewTemplateService(ctx)
	if err != nil {
		helper.ErrorfLogger(ctx, "NewTemplateService err%+v", err)
		return
	}
	reply, err := tempService.GetReply(ctx, msg)
	if err != nil {
		helper.ErrorfLogger(ctx, "GetReply err%+v", err)
		return
	}

	if reply == nil {
		return
	}

	messageConfig := tgbotapi.NewMessage(chatID, *reply)
	_, err = bot.Send(messageConfig)
	if err != nil {
		helper.ErrorfLogger(ctx, "send err%+v", err)
		return
	}
	msg.ReplyText = reply
	msg.ReplyTime = time.Now()
	msg.UpdateTime = time.Now()
	_, err = model.UpdateTgMessage(ctx, msg)
	if err != nil {
		helper.ErrorfLogger(ctx, "update message error%+v", err)
		return
	}
}

func StartReceiveTgMessage() {
	go helper.DealMessage(handleMessage)
}
