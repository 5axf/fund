package ginResponseUtil

/**
gin 请求数据返回封装
使用方法：
utli := ginUtil.Gin{c}
utli.Response(-1,"no data",nil)

*/
import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(200, response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}


func (g *Gin) Ok() {
	g.Ctx.JSON(200, response{
		Code:    200,
		Message: "操作成功",
		Data:    nil,
	})
	return
}
func (g *Gin) OkMsg(msg string) {
	g.Ctx.JSON(200, response{
		Code:    200,
		Message: msg,
		Data:    nil,
	})
	return
}
func (g *Gin) OkData(data interface{}) {
	g.Ctx.JSON(200, response{
		Code:    200,
		Message: "操作成功",
		Data:    data,
	})
	return
}
func (g *Gin) OkMsgData(msg string, data interface{}) {
	g.Ctx.JSON(200, response{
		Code:    200,
		Message: msg,
		Data:    data,
	})
	return
}


func (g *Gin) Error() {
	g.Ctx.JSON(200, response{
		Code:    500,
		Message: "操作失败",
		Data:    nil,
	})
	return
}
func (g *Gin) ErrorMsg(msg string) {
	g.Ctx.JSON(200, response{
		Code:    500,
		Message: msg,
		Data:    nil,
	})
	return
}
func (g *Gin) ErrorData(data interface{}) {
	g.Ctx.JSON(200, response{
		Code:    500,
		Message: "操作失败",
		Data:    data,
	})
	return
}
func (g *Gin) ErrorMsgData(msg string, data interface{}) {
	g.Ctx.JSON(200, response{
		Code:    500,
		Message: msg,
		Data:    data,
	})
	return
}