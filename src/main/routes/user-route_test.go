package routes

import (
	"bytes"
	"encoding/json"
	"github.com/douglasdennys/go-mongodb/src/domain/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func request(method, path string, e *echo.Echo) (int, string) {
	dataJSON := test.MockAddUserParam()
	convert, _ := json.Marshal(dataJSON)
	req := httptest.NewRequest(method, path, bytes.NewReader(convert))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}

func testMethod(t *testing.T, method, path string, e *echo.Echo) {
	p := reflect.ValueOf(path)
	h := reflect.ValueOf(func(c echo.Context) error {
		return c.String(http.StatusCreated, method)
	})
	i := interface{}(e)
	reflect.ValueOf(i).MethodByName(method).Call([]reflect.Value{p, h})
	code, body := request(method, path, e)
	assert.Equal(t, method, body)
	assert.Equal(t, 201, code)
}

func TestUserRouter(t *testing.T) {
	e := echo.New()
	testMethod(t, http.MethodPost, "/", e)
}
