package routes

import (
	"github.com/kataras/iris/v12/context"
	"rogxd/rogxd/bus/controllers"
)

type (
	Context struct {
		Base *context.Context
		RogX *controllers.RogXConn
	}
	SuccessResponse struct {
		Message string `json:"message"`
	}
	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	SuccessResponseWithData struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func (ctx Context) Query() *context.RequestParams {
	return (*ctx.Base).Params()
}

func (ctx Context) Forbid(message string) {
	ctx.Error(403, message)
}

func (ctx Context) Bad(message string) {
	ctx.Error(400, message)
}

func (ctx Context) Json(data interface{}) {
	ctx.JsonWithCode(200, data)
}

func (ctx Context) Body(out interface{}) error {
	return (*ctx.Base).ReadJSON(&out)
}

func (ctx Context) JsonWithCode(code int, data interface{}) {
	base := *ctx.Base
	base.StatusCode(code)
	base.JSON(data)
}

func (ctx Context) Success(message string, data ...interface{}) {
	if len(data) > 0 {
		ctx.Json(&SuccessResponseWithData{
			Message: message,
			Data:    data[0],
		})
	} else {
		ctx.Json(&SuccessResponse{
			Message: message,
		})
	}
}

func (ctx Context) Error(code int, message string) {
	ctx.JsonWithCode(code, &ErrorResponse{
		Code:    code,
		Message: message,
	})
}
