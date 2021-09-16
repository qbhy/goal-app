package exceptions

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/events"
	"github.com/qbhy/goal/http"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/utils"
	"reflect"
)

var (
	dontReportExceptions []reflect.Type
)

func init() {
	dontReportExceptions = utils.ConvertToTypes(
	// events.EventException{}, // 不上报的异常
	)
}

type ExceptionHandler struct {
}

func (handler ExceptionHandler) ShouldReport(exception contracts.Exception) bool {
	for _, e := range dontReportExceptions {
		if utils.IsSameStruct(e, exception) {
			return false
		}
	}
	return true
}

func (handler ExceptionHandler) Report(exception contracts.Exception) {
	// todo: 上报异常
}

func (handler ExceptionHandler) Handle(exception contracts.Exception) {
	if handler.ShouldReport(exception) {
		handler.Report(exception)
	}

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
