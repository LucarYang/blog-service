package api

import (
	"blog-service/global"
	"blog-service/internal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/convert"
	"blog-service/pkg/errcode"
	"blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// @Summary 图片上传
// @Description
// @Tags file
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Param type formData int true "type" Enums(0, 1,2,3,4,5,6,7) default(1)
// @Produce  json
// @Success 200 {object} app.ResponseHandle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	//log.WithFields(log.Fields{"":"c"}).Info(c)
	//log.WithFields(log.Fields{"":"file"}).Info(file)
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	//log.WithFields(log.Fields{"":"fileType"}).Info(fileType)
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c,"c", "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
