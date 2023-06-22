package api_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/injection"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	os.Setenv("DOMAIN_CONFIG", "../domain.config.json")
	os.Setenv("GIN_LOG", "../gin.log")
	app, err := injection.ProvideApplication()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"account": "account3", "password": "e10adc3949ba59abbe56e057f20f883e"}`))
	req, _ := http.NewRequest("POST", "/login", body)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	fmt.Println(recorder.Body.String())
	cookies := resp.Cookies()
	for _, v := range cookies {
		fmt.Println(v.String())
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestLogout(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	os.Setenv("DOMAIN_CONFIG", "../domain.config.json")
	os.Setenv("GIN_LOG", "../gin.log")
	app, err := injection.ProvideApplication()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/logout", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "2"}
	token_cookie := &http.Cookie{Name: "token", Value: "20502c896d3d2a4322fe36c0a6a43a0a"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	fmt.Println(recorder.Body.String())
	assert.Equal(t, 200, recorder.Code)
}
