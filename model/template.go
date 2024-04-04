package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"messageBot/helper"
	"time"
)

const TableNameTemplate = "template"

type Template struct {
	ID         int64     `gorm:"column:id; PRIMARY_KEY"`
	Command    string    `gorm:"column:command"` // Comment: 正则表达式
	Reg        string    `gorm:"column:reg"`
	Reply      string    `gorm:"column:reply"`       // Default: CURRENT_TIMESTAMP
	CreateTime time.Time `gorm:"column:create_time"` // Default: 0000-00-00 00:00:00
	UpdateTime time.Time `gorm:"column:update_time"`
}

func InsertTemplate(ctx *gin.Context, d Template) (com *Template, err error) {
	db := helper.MysqlClient
	err = db.SetCtx(ctx).Create(&d).Error
	return &d, err
}

func DeleteTemplate(ctx *gin.Context, id int) (effectedRow int64, err error) {
	db := helper.MysqlClient
	res := db.SetCtx(ctx).Table(TableNameTemplate).Delete(&Template{}, "id = ? ", id)
	return res.RowsAffected, res.Error
}

func GetTemplate(ctx *gin.Context, id int) (t *Template, err error) {
	t = new(Template)
	db := helper.MysqlClient
	err = db.SetCtx(ctx).Table(TableNameTemplate).Where("id = ?", id).First(&t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return t, err
}

func CountTemplateList(ctx *gin.Context, name string) (count int, err error) {
	db := helper.MysqlClient
	database := db.SetCtx(ctx)
	database = database.Table(TableNameTemplate).Select("count(*) as total")
	if len(name) > 0 {
		database = database.Where("name like ?", "%"+name+"%")
	}
	err = database.Count(&count).Error
	return count, err
}

func GetTemplateList(ctx *gin.Context, limit int, offset int) (TemplateList []Template, err error) {
	db := helper.MysqlClient
	database := db.SetCtx(ctx).Table(TableNameTemplate)
	err = database.Order("create_time desc").Limit(limit).Offset(offset).Find(&TemplateList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return TemplateList, err
}

func GetAllTemplateList(ctx *gin.Context) (TemplateList []Template, err error) {
	db := helper.MysqlClient
	database := db.SetCtx(ctx).Table(TableNameTemplate)
	err = database.Order("create_time desc").Find(&TemplateList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return TemplateList, err
}

func UpdateTemplate(ctx *gin.Context, d Template) (effectedRow int64, err error) {
	db := helper.MysqlClient
	res := db.SetCtx(ctx).Table(TableNameTemplate).Where("id = ?", d.ID).Updates(d)
	return res.RowsAffected, res.Error
}
