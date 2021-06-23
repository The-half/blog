package service

import (
	"blog/internal/dao"
	"blog/internal/model"
	"blog/pkg/app"
)

type Article struct {
	ID uint32 `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
	Tag *model.Tag `json:"tag"`
}

//统计状态
type CountArticleRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//对文章进行查询
type ArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state" binding:"oneof=0 1"`
}

//对所有列表进行查询
type ArticleListRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state, default=1" binding:"oneof=0 1"`
	TagID uint32 `form:"tag_id" bindint:"required,gte=1"`
}

//创建一篇文章
type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//更新一个文章的信息
type UpdateArticleRequest struct{
	ID uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

//删除一个文章
type DeletedArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//获取单篇文章
func (svc *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := svc.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	articleTag, err := svc.dao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}

	tag, err := svc.dao.GetTag(articleTag.TagID,param.State)
	if err != nil {
		return nil, err
	}

	return &Article{
		ID : article.ID,
		Title: article.Title,
		Desc: article.Desc,
		Content: article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State: article.State,
		Tag: &tag,
	}, nil
}

//获取文章列表
func (svc *Service) GetArticleList(param *ArticleListRequest, paper *app.Pager) ([]*Article, int, error) {
	articleCount, err := svc.dao.CountArticleListByTagID(param.TagID,param.State)
	if err != nil {
		return nil, 0, err
	}
	articles, err := svc.dao.GetArticleListByTagID(param.TagID, param.State,paper.Page,paper.PageSize)
	if err != nil {
		return nil, 0, err
	}
	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			ID : article.ArticleID,
			Title : article.ArticleTitle,
			Desc : article.ArticleDesc,
			Content : article.ArticleDesc,
			CoverImageUrl: article.CoverImageUrl,
			Tag : &model.Tag{
				Model: &model.Model{
					ID: article.TagID,
				},
				Name: article.TagName,
			},
		})
	}
	return articleList, articleCount, nil
}

//创建一篇文章
func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := svc.dao.CreateArticle(&dao.Article{
		Title: param.Title,
		Desc: param.Desc,
		Content: param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State: param.State,
		CreatedBy: param.CreatedBy,
	})
	if err != nil {
		return err
	}
	err = svc.dao.CreateArticleTag(article.ID, param.TagID, param.CreatedBy)
	if err != nil {
		return err
	}
	return nil
}

//更新一篇文章
func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	err := svc.dao.UpdateArticle(&dao.Article{
		ID: param.ID,
		Title: param.Title,
		Desc: param.Desc,
		Content: param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State: param.State,
		ModifiedBy: param.ModifiedBy,
	})
	if err != nil {
		return err
	}
	err = svc.dao.UpdateArticleTag(param.ID, param.TagID, param.ModifiedBy)
	if err != nil {
		return err
	}
	return nil
}


//删除一篇文章
func (svc *Service) DeleteArticle(param *DeletedArticleRequest) error {
	err := svc.dao.DeleteArticle(param.ID)
	if err != nil {
		return err
	}
	err = svc.dao.DeleteArticleTag(param.ID)
	if err != nil {
		return err
	}
	return nil
}