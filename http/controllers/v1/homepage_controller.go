package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/logs"
	"goalapp/http/requests"
)

func Homepage(context echo.Context) interface{} {
	var request requests.LoginRequest
	logs.WithError(context.Bind(request)).Debug("绑定参数失败")
	request.Assure()

	return "hello goal!"
}
