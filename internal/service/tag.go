package service

import (
	"blog/internal/model"
	"blog/pkg/app"
)

//统计状态
type CountTagRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//对所有列表进行查询
type TagListRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//创建一个标签
type CreateTagRequest struct {
	Name string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//更新一个标签的信息
type UpdateTagRequest struct{
	ID uint32 `form:"id" binding:"required,gte=1"`
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

//删除一个标签
type DeletedTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func(svc *Service) GetTagList(param *TagListRequest, paper *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name,param.State,paper.Page,paper.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeletedTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}

