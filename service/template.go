package service

import (
	"github.com/gin-gonic/gin"
	"messageBot/helper"
	"messageBot/model"
	"regexp"
)

var templateCache *TemplateCache

type TemplateCache struct {
	List     []*model.Template
	Commands map[int64]*model.Template
	Regs     map[int64]*model.Template
}

type TemplateService struct {
}

func NewTemplateService(ctx *gin.Context) (*TemplateService, error) {
	tpl := new(TemplateService)
	if templateCache == nil {
		err := iniTemplates(ctx)
		if err != nil {
			return nil, err
		}
	}
	return tpl, nil
}

func iniTemplates(ctx *gin.Context) error {
	templates, err := model.GetAllTemplateList(ctx)
	if err != nil {
		helper.ErrorfLogger(ctx, "init error %+v", err)
		return err
	}

	cache := TemplateCache{
		List:     make([]*model.Template, 0, len(templates)),
		Commands: make(map[int64]*model.Template),
		Regs:     make(map[int64]*model.Template),
	}

	for _, tpl := range templates {
		newT := tpl
		if tpl.Command != "" {
			cache.Commands[tpl.ID] = &newT
		}

		if tpl.Reg != "" {
			cache.Regs[tpl.ID] = &newT
		}
		cache.List = append(cache.List, &newT)
	}
	templateCache = &cache
	return nil
}

/**
 * 创建
 */
func (*TemplateService) Create(ctx *gin.Context, tpl model.Template) (template *model.Template, err error) {

	if tpl.Reg != "" {

		_, err := regexp.Compile(tpl.Reg)

		if err != nil {
			helper.ErrorfLogger(ctx, "reg error %+v", err)
			return nil, err
		}
	}
	newTpl, err := model.InsertTemplate(ctx, tpl)
	if err != nil {
		return nil, err
	}

	templateCache = nil

	return newTpl, nil
}

/**
 * 获取
 */
func (*TemplateService) GetById(ctx *gin.Context, id int) (com *model.Template, err error) {
	return model.GetTemplate(ctx, id)
}

/**
 * 删除
 */
func (*TemplateService) Remove(ctx *gin.Context, id int) error {
	_, err := model.DeleteTemplate(ctx, id)
	if err != nil {
		helper.ErrorfLogger(ctx, "delete error %+v", err)
		return err
	}
	templateCache = nil
	return nil
}

/**
 * 更新
 */
func (*TemplateService) Update(ctx *gin.Context, tpl model.Template) error {
	_, err := model.UpdateTemplate(ctx, tpl)
	if err != nil {
		return err
	}
	templateCache = nil
	return nil
}

/**
 * 获取分页信息
 */
func (*TemplateService) GetPageInfo(ctx *gin.Context, pageSize int) (totalCount int, totalPage int, err error) {
	totalCount = len(templateCache.List)
	totalPage = totalCount / pageSize
	if totalCount%pageSize != 0 {
		totalPage += 1
	}
	return totalCount, totalPage, nil
}

/**
 * 获取列表
 */
func (*TemplateService) GetList(ctx *gin.Context, page int, pageSize int) (tplList []*model.Template, err error) {
	offset := (page - 1) * pageSize
	if offset > (len(templateCache.List) - 1) {
		return nil, nil
	}
	endSet := offset + pageSize
	if endSet >= len(templateCache.List) {
		endSet = len(templateCache.List)
	}
	newList := templateCache.List[offset:endSet]
	tplList = make([]*model.Template, len(newList))
	copy(tplList, newList)
	return tplList, nil
}

func (*TemplateService) GetReply(ctx *gin.Context, msg model.TgMessage) (reply *string, err error) {
	if msg.IsCommand {
		for _, tpl := range templateCache.Commands {
			if tpl.Command == msg.Command {
				return &tpl.Reply, nil
			}
		}
		return nil, nil
	}

	for _, tpl := range templateCache.Regs {
		found, err := regexp.MatchString(tpl.Reg, *msg.TextContent)

		if err != nil {
			helper.ErrorfLogger(ctx, "match error %+v", err)
		}

		if found {
			return &tpl.Reply, nil
		} //todo 可以用模板
	}
	return nil, nil
}
