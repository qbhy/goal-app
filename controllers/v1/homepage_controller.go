package v1

import (
	"github.com/labstack/echo/v4"
)

func Homepage(context echo.Context) interface{} {
	return "hello goal!"
}
