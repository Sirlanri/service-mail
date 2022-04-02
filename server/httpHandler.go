package server

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

//NotFound -handler 前端请求地址错误，调用此handler处理
func NotFound(ctx iris.Context) {
	fmt.Println("404-找不到此路由/路径:", ctx.RequestPath(true))
	ctx.StatusCode(404)
	ctx.WriteString("路由/请求地址错误")
}

//接受邮件请求的http handler
func SendHandler(con iris.Context) {
	var res SendData
	err := con.ReadJSON(&res)
	if err != nil {
		fmt.Println("传入格式有误", err.Error())
		con.StatusCode(iris.StatusForbidden)
		return
	}
	err = Sendmail(res.Addr, res.Topic, res.Content)
	if err != nil {
		fmt.Println("发送失败！", err.Error())
		con.StatusCode(iris.StatusForbidden)
		return
	}
}

//使用http接收的post请求
type SendData struct {
	Addr    string `json:"addr"`
	Topic   string `json:"topic"`
	Content string `json:"content"`
}
