package testingutils

import (
	"github.com/labstack/echo"
	"github.com/google/uuid"
	"net/http/httptest"
	"encoding/json"
	"strings"
	"io"
	"bytes"
	"net/http"
	"boardbots/server/context"
	"boardbots/quoridor"
	"reflect"
)

var TestUUID = uuid.MustParse("c67a791f-1d1b-41ae-b21b-14f79d4fdf66")
var TestMissingUUID = uuid.MustParse("df834d72-8d96-4010-bb37-36c60d9309cd")
type header struct {
name, value string
}
func GetPayloadFromResult(httpError *echo.HTTPError, response *interface{}) interface{} {
	return httpError.Message.(interface{})
}

func ToJson(body interface{}) io.Reader {
	b, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(b)
}

type FakeContextBuilder struct {
	Path, Payload, Method string
	Headers               []header
	Player                context.PlayerPrinciple
	Game                  quoridor.Game
}

func DefaultFakeContextBuilder() FakeContextBuilder {
	return FakeContextBuilder{
		Payload: "",
		Method:  http.MethodPost,
		Path:    "/defaulttest",
		Headers: make([]header, 0, 0),
		Player:  context.PlayerPrinciple{},
	}
}
func (b FakeContextBuilder) Override(override FakeContextBuilder) FakeContextBuilder {
	empty := FakeContextBuilder{}
	baseRef := reflect.ValueOf(b)
	emptyRef := reflect.ValueOf(&empty).Elem()
	overrideRef := reflect.ValueOf(override)
	fcbType := emptyRef.Type()
	for fieldIdx := 0; fieldIdx < fcbType.NumField(); fieldIdx++ {
		field := fcbType.Field(fieldIdx)
		overrideValue := overrideRef.FieldByName(field.Name)
		if isEmpty(field, overrideValue) {
			emptyRef.FieldByName(field.Name).Set(overrideValue)
		} else {
			baseVal := baseRef.FieldByName(field.Name)
			emptyRef.FieldByName(field.Name).Set(baseVal)
		}
	}
	return empty
}

func isEmpty(field reflect.StructField, overrideValue reflect.Value) bool {
	return field.Type.Comparable() && overrideValue.Interface() != reflect.Zero(field.Type).Interface() ||
		strings.HasPrefix(field.Type.Name(), "[]") && overrideValue.Len() == 0
}

func Build(builder FakeContextBuilder) (context.DefaultBBContext, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(builder.Method, builder.Path, strings.NewReader(builder.Payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	for _, header := range builder.Headers {
		req.Header.Set(header.name, header.value)
	}
	recorder := httptest.NewRecorder()
	echoContext := e.NewContext(req, recorder)
	bbCtx := context.DefaultBBContext{
		Context :        echoContext,
		PlayerPrinciple: builder.Player,
		Game :           &builder.Game,
	}
	return bbCtx, recorder
}
/**
 *
 */
func ReadBodyFromRecorder(payload *httptest.ResponseRecorder, response interface{}){
	json.NewDecoder(payload.Body).Decode(response)
}

func FakeContext(method, path, payload string) (echo.Context, *httptest.ResponseRecorder){
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	return e.NewContext(req, recorder), recorder
}

func FakeBBContext(path, payload string) (context.DefaultBBContext, *httptest.ResponseRecorder) {
	ctx, rec := FakeContext(http.MethodPost, path, payload)
	bbCtx := context.DefaultBBContext{}
	bbCtx.Context = ctx
	return bbCtx, rec
}