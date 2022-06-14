package app

import (
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
响应处理
*/

type Response struct {
	Ctx *gin.Context
}

type ResponseHandle struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Error string `json:"error"`
	Data  interface{}
}
type ListData struct {
	List   interface{}
	Pagers Pager
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	rh := ResponseHandle{}
	rh.Data = data
	rh.Code = 0
	rh.Msg = "success"
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code": errcode.Success.Code(),
		"msg":  errcode.Success.Msg(),
		"data": data,
	})
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	list = ListData{
		List: list,
		Pagers: Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	}
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code": errcode.Success.Code(),
		"msg":  errcode.Success.Msg(),
		"data": list,
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
