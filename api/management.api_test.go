package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/injection"
	"github.com/stretchr/testify/assert"
)

func TestGetRoleList(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	os.Setenv("DOMAIN_CONFIG", "../domain.config.json")
	os.Setenv("GIN_LOG", "../gin.log")
	app, err := injection.ProvideApplication()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	// body := bytes.NewReader([]byte(`{"current": 3, "pageSize": 1, "order": false, "sort": "id"}`))
	query := "?current=1&pagesize=4&sort=id&order=true&like=%E7%AE%A1%E7%90%86%E5%91%98"
	req, _ := http.NewRequest(http.MethodGet, "/role/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "7b3cf626947056ff98ae7e41e707367a"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}
