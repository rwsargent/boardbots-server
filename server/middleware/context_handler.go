package middleware

import (
	"github.com/labstack/echo"
	"boardbots/server/context"
)

func ContextHander(h echo.HandlerFunc) echo.HandlerFunc {
	return func (ctx echo.Context) error {
		customCtx := &context.DefaultBBContext{Context :ctx}
		return h(customCtx)
	}
}


