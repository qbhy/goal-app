package exceptions

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/events"
	"github.com/qbhy/goal/http"
	"github.com/qbhy/goal/logs"
)

type ExceptionHandler struct{}

func (handler ExceptionHandler) Handle(exception contracts.Exception) {
	switch e := exception.(type) {
	case events.EventException:
		logs.WithException(e).Info("事件报错啦")
	case http.HttpException:
		logs.WithException(e).Error("控制器报错啦")
		_ = e.Context.String(500, e.Error())
	default:
		logs.WithException(e).Info("默认异常")
	}
}
