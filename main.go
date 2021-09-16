package main

import (
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/events"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/http/routes"
	"github.com/qbhy/goal/logs"
	appExceptions "goalapp/exceptions"
)

func main() {
	// 注册异常处理器
	exceptions.SetExceptionHandler(appExceptions.ExceptionHandler{})

	// 注册事件监听器
	events.SetEventListeners(map[contracts.EventName][]contracts.EventListener{})
	router := routes.New()

	router.Get("/", func(context echo.Context) interface{} {
		return "返回了啥"
	})

	group := router.Group("/user")

	group.Get("/test", func(context echo.Context) interface{} {
		return "测试"
	})

	logs.WithError(router.Start(":8000")).Debug("服务器报错了")
}
