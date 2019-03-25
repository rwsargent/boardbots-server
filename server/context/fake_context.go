package context

import (
	"boardbots/quoridor"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
)

type header struct {
	name, value string
}

type FakeContextBuilder struct {
	Path, Payload, Method string
	Headers               []header
	Player                PlayerPrinciple
	Game                  quoridor.Game
}

func DefaultFakeContextBuilder() FakeContextBuilder {
	return FakeContextBuilder{
		Payload: "",
		Method:  http.MethodPost,
		Path:    "/defaulttest",
		Headers: make([]header, 0, 0),
		Player:  PlayerPrinciple{},
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

func Build(builder FakeContextBuilder) (DefaultBBContext, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(builder.Method, builder.Path, strings.NewReader(builder.Payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	for _, header := range builder.Headers {
		req.Header.Set(header.name, header.value)
	}
	recorder := httptest.NewRecorder()
	echoContext := e.NewContext(req, recorder)
	bbCtx := DefaultBBContext{
		Context:         echoContext,
		PlayerPrinciple: builder.Player,
		Game:            &builder.Game,
	}
	return bbCtx, recorder
}

func ReadBodyFromRecorder(payload *httptest.ResponseRecorder, response interface{}) {
	json.NewDecoder(payload.Body).Decode(response)
}

func FakeContext(method, path, payload string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	return e.NewContext(req, recorder), recorder
}

func FakeBBContext(path, payload string) (DefaultBBContext, *httptest.ResponseRecorder) {
	ctx, rec := FakeContext(http.MethodPost, path, payload)
	bbCtx := DefaultBBContext{}
	bbCtx.Context = ctx
	return bbCtx, rec
}
