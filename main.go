package main

import (
	"github.com/Sirlanri/service-mail/server"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, server.NotFound)
	mail := app.Party("/mail").AllowMethods(iris.MethodOptions)
	{
		mail.Post("/send")
	}
}
