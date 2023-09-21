package api_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/injection"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CreateTestApp() (*gin.Engine, error) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	os.Setenv("DOMAIN_CONFIG", "../domain.config.json")
	os.Setenv("GIN_LOG", "../gin.log")
	app, err := injection.ProvideApplication()
	if err != nil {
		return nil, err
	}
	return app, nil
}

func TestGetRoleList(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?current=2&pagesize=5"
	req, _ := http.NewRequest(http.MethodGet, "/role/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "ed51c2f46c751f31b34c01b0234b9602"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestAddRole(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`[{"name": "new role8", "description": "oh my girl"}, {"name": "new role111", "description": "stayc"}]`))
	// body := bytes.NewReader([]byte(`{"name": ""}`))
	req, _ := http.NewRequest(http.MethodPost, "/role/add", body)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "ed51c2f46c751f31b34c01b0234b9602"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestModifyARole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	os.Setenv("DOMAIN_CONFIG", "../domain.config.json")
	os.Setenv("GIN_LOG", "../gin.log")
	app, err := injection.ProvideApplication()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"id": 5, "description": "something else", "isDisabled": false}`))
	req, _ := http.NewRequest(http.MethodPost, "/role/modify", body)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "0b3ae0f9c2a633260c8e80b2c180e70a"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetARoleById(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/role/get/9", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "ed51c2f46c751f31b34c01b0234b9602"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestRemoveRolesByIds(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	os.Setenv("DOMAIN_CONFIG", "../domain.config.json")
	os.Setenv("GIN_LOG", "../gin.log")
	app, err := injection.ProvideApplication()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"ids":[4, 7, 9]}`))
	req, _ := http.NewRequest(http.MethodPost, "/role/remove", body)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "d8d519f1c826411bee4cb8e9eedb9734"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetPermssionList(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	os.Setenv("DOMAIN_CONFIG", "../domain.config.json")
	os.Setenv("GIN_LOG", "../gin.log")
	app, err := injection.ProvideApplication()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := ""
	req, _ := http.NewRequest(http.MethodGet, "/permission/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "ed51c2f46c751f31b34c01b0234b9602"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestAddAPermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	os.Setenv("DOMAIN_CONFIG", "../domain.config.json")
	os.Setenv("GIN_LOG", "../gin.log")
	app, err := injection.ProvideApplication()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`[{"action": "action15", "description": "test add action15"}, {"action": "action16", "description": ""}, {"action": "action17", "description": "test add action11115"}]`))
	// body := bytes.NewReader([]byte(`{"name": ""}`))
	req, _ := http.NewRequest(http.MethodPost, "/permission/add", body)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "4933ddad9c76c8657cbd4ae5fd44339f"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestModifyAPermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	os.Setenv("DOMAIN_CONFIG", "../domain.config.json")
	os.Setenv("GIN_LOG", "../gin.log")
	app, err := injection.ProvideApplication()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"id": 10, "action": "action10", "isDisabled": true}`))
	req, _ := http.NewRequest(http.MethodPost, "/permission/modify", body)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "549f909f1108f1694e00e46d5ca514b1"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetAPermissionById(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/permission/get/12", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "ed51c2f46c751f31b34c01b0234b9602"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestRemovePermissionsByIds(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"ids":[11, 12]}`))
	req, _ := http.NewRequest(http.MethodPost, "/permission/remove", body)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "549f909f1108f1694e00e46d5ca514b1"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetPermissionsByRoleId(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/role/3/get/permissions", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "ed51c2f46c751f31b34c01b0234b9602"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestModifyPermissionsForRoles(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"rids": [2], "pids": [1, 2]}`))
	req, _ := http.NewRequest(http.MethodPost, "/role/modify/permissions", body)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "ef7f90af5c30a440467df6b704da36f2"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetUserList(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?current=-991&sort=id&order=true"
	req, _ := http.NewRequest(http.MethodGet, "/user/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "ed51c2f46c751f31b34c01b0234b9602"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetAllPermission(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/permission/getall", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "8483f6a8aaed3ce434c5d3fb51aa26bd"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}
