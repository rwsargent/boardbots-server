package middleware

import (
	"boardbots-server/server/context"
	"github.com/labstack/echo"
)

func ContextHander(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		customCtx := &context.DefaultBBContext{Context: ctx}
		return h(customCtx)
	}
}
