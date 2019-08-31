package transport

import (
	"github.com/labstack/echo"
)

type (
	HandledError struct {
		Code    int
		Message string
	}
)

func NewHandledError(code int, message string) HandledError {
	return HandledError{
		Code:    code,
		Message: message,
	}
}
func (e HandledError) Error() string {
	return e.Message
}

func EchoErrorHandler(server *echo.Echo) {
	server.HTTPErrorHandler = func(e error, context echo.Context) {
		if handledError, ok := e.(HandledError); ok {
			resp := BaseResponse{
				Error: handledError.Error(),
			}
			if !context.Response().Committed {
				err := context.JSON(handledError.Code, resp)
				if err != nil {
					context.Logger().Error(err)
				}
			}
		} else {
			server.DefaultHTTPErrorHandler(e, context)
		}
	}
}
