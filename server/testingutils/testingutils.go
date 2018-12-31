package testingutils

import (
	"github.com/labstack/echo"
	"github.com/google/uuid"
	"net/http/httptest"
	"encoding/json"
	"strings"
	"io"
	"bytes"
)

var TestUUID = uuid.MustParse("c67a791f-1d1b-41ae-b21b-14f79d4fdf66")
var TestMissingUUID = uuid.MustParse("df834d72-8d96-4010-bb37-36c60d9309cd")

func getPayloadFromResult(httpError *echo.HTTPError, response *interface{}) interface{} {
	return httpError.Message.(interface{})
}

func ToJson(body interface{}) io.Reader {
	b, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(b)
}


/**
 *
 */
func FillResponseFromPayload(payload *httptest.ResponseRecorder, response interface{}){
	json.NewDecoder(payload.Body).Decode(response)
}

func FakeContext(method, path, payload string) (echo.Context, *httptest.ResponseRecorder){
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	return e.NewContext(req, recorder), recorder
}