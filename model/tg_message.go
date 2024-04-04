package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"messageBot/helper"
	"time"
)

const TableNameTgMessage = "tg_message"

type TgMessage struct {
	ID          int64     `gorm:"column:id; PRIMARY_KEY"`
	ChatID      int64     `gorm:"column:chat_id"`      // Comment: 用户id
	UID         int64     `gorm:"column:u_id"`         // Comment: 是否是命令
	IsCommand   bool      `gorm:"column:is_command"`   // Comment: 命令类型
	Command     string    `gorm:"column:command"`      // Comment: 文本内容
	TextContent *string   `gorm:"column:text_content"` // Default: CURRENT_TIMESTAMP
	ReceiveTime time.Time `gorm:"column:receive_time"` // Comment: 消息发送时间 // Default: 0000-00-00 00:00:00
	SendTime    time.Time `gorm:"column:send_time"`
	TemplateID  int       `gorm:"column:template_id"` // Default: 0000-00-00 00:00:00
	ReplyTime   time.Time `gorm:"column:reply_time"`
	ReplyText   *string   `gorm:"column:reply_text"`  // Default: CURRENT_TIMESTAMP
	CreateTime  time.Time `gorm:"column:create_time"` // Default: 0000-00-00 00:00:00
	UpdateTime  time.Time `gorm:"column:update_time"`
}

func InsertTgMessage(ctx *gin.Context, m TgMessage) (message *TgMessage, err error) {
	db := helper.MysqlClient
	err = db.SetCtx(ctx).Create(&m).Error
	return &m, err
}

func GetTgMessage(ctx *gin.Context, id int) (m *TgMessage, err error) {
	m = new(TgMessage)
	db := helper.MysqlClient
	err = db.SetCtx(ctx).Table(TableNameTgMessage).Where("id = ?", id).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return m, err
}

func CountTgMessageList(ctx *gin.Context) (count int, err error) {
	db := helper.MysqlClient
	database := db.SetCtx(ctx)
	database = database.Table(TableNameTgMessage).Select("count(*) as total")
	err = database.Count(&count).Error
	return count, err
}

func GetTgMessageList(ctx *gin.Context, limit int, offset int) (TgMessageList []TgMessage, err error) {
	db := helper.MysqlClient
	database := db.SetCtx(ctx).Table(TableNameTgMessage)

	err = database.Order("create_time desc").Limit(limit).Offset(offset).Find(&TgMessageList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return TgMessageList, err
}

func UpdateTgMessage(ctx *gin.Context, d TgMessage) (effectedRow int64, err error) {
	db := helper.MysqlClient
	res := db.SetCtx(ctx).Table(TableNameTgMessage).Where("id = ?", d.ID).Updates(d)
	return res.RowsAffected, res.Error
}
