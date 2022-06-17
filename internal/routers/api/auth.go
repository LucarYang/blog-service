package api

import (
	"blog-service/global"
	"blog-service/internal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 新增标签
// @Tags auth
// @Produce  json
// @Param body body service.AuthRequest true "入参"
// @Success 200 {object} app.ResponseHandle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	zap.L().Debug("this is hello func", zap.String("user", "name"), zap.Int("age", 1))
	fmt.Println("zap 后……")
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c,"app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c,"svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c,"app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
