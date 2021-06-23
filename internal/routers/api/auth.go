package api

import (
	"blog/global"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
			global.Logger.Errorf(c, "app.bindAndValid errs: %v", errs)
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c,"app.GenerateToken errs:%v",err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
