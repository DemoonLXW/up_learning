package api_test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
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

func TestImportSchool(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	f, err := os.Open("../importSchools2.xlsx")
	assert.Nil(t, err)
	defer f.Close()

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("import", filepath.Base(f.Name()))
	assert.Nil(t, err)

	_, err = io.Copy(part, f)
	assert.Nil(t, err)
	err = writer.Close()
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/school/import", &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "e395616bf00310ee678ec35bc0876943"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetSchoolList(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?current=1&pagesize=3&sort=id&order=true"
	req, _ := http.NewRequest(http.MethodGet, "/school/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "c7a18ec1c1c94696059d629d47899262"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetSampleOfImport(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/file/getsample?type=import teacher", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "917f6cd13b307b54c46e341b4c4e24ba"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	disposition := recorder.Header().Get("Content-Disposition")
	fmt.Println(disposition)
	re := regexp.MustCompile(`filename="(.+)"`)
	res := re.FindStringSubmatch(disposition)
	assert.NotNil(t, res)
	fmt.Println(res[1])

	down, err := os.Create("temp/download/" + res[1])
	assert.Nil(t, err)

	_, err = recorder.Body.WriteTo(down)
	assert.Nil(t, err)
	// fmt.Println(recorder.Body.String())
}

func TestGetStudentList(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?pagesize=1&sort=id&order=true&like=ly"
	req, _ := http.NewRequest(http.MethodGet, "/student/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "3efd0f809c9dc5d317369d38a374201c"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestImportCollege(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	f, err := os.Open("../import.xlsx")
	assert.Nil(t, err)
	defer f.Close()

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("import", filepath.Base(f.Name()))
	assert.Nil(t, err)

	_, err = io.Copy(part, f)
	assert.Nil(t, err)
	err = writer.Close()
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/college/import", &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "3efd0f809c9dc5d317369d38a374201c"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetCollegeList(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?pagesize=1&sort=id&order=true&like=1"
	req, _ := http.NewRequest(http.MethodGet, "/college/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "4a91f00f7d3dbd6a2a7c078a9e7ce62b"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestImportMajorByCollegeID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	f, err := os.Open("../import.xlsx")
	assert.Nil(t, err)
	defer f.Close()

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("import", filepath.Base(f.Name()))
	assert.Nil(t, err)
	_, err = io.Copy(part, f)
	assert.Nil(t, err)

	err = writer.WriteField("id", "4")
	assert.Nil(t, err)

	err = writer.Close()
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/major/import", &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "70a076f5888b9faffe576891656fe6cc"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetMajorList(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?pagesize=1&sort=id"
	req, _ := http.NewRequest(http.MethodGet, "/major/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "70a076f5888b9faffe576891656fe6cc"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestImportClassByMajorID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	f, err := os.Open("../import.xlsx")
	assert.Nil(t, err)
	defer f.Close()

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("import", filepath.Base(f.Name()))
	assert.Nil(t, err)
	_, err = io.Copy(part, f)
	assert.Nil(t, err)

	err = writer.WriteField("id", "6")
	assert.Nil(t, err)

	err = writer.Close()
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/class/import", &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "d0082b284c33f13c2cdea0a76c942caf"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetClassList(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?current=2&pagesize=1&sort=id&like=class1"
	req, _ := http.NewRequest(http.MethodGet, "/class/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "d0082b284c33f13c2cdea0a76c942caf"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestImportStudentByClassID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	f, err := os.Open("../import.xlsx")
	assert.Nil(t, err)
	defer f.Close()

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("import", filepath.Base(f.Name()))
	assert.Nil(t, err)
	_, err = io.Copy(part, f)
	assert.Nil(t, err)

	err = writer.WriteField("id", "1")
	assert.Nil(t, err)

	err = writer.Close()
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/student/import", &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "f87fb5fd46a93314cf9b46d0f253a838"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetTeacherList(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?current=2&pagesize=2&sort=id&order=true"
	req, _ := http.NewRequest(http.MethodGet, "/teacher/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "917f6cd13b307b54c46e341b4c4e24ba"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestImportTeacherByCollegeID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	f, err := os.Open("../import.xlsx")
	assert.Nil(t, err)
	defer f.Close()

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("import", filepath.Base(f.Name()))
	assert.Nil(t, err)
	_, err = io.Copy(part, f)
	assert.Nil(t, err)

	err = writer.WriteField("id", "2")
	assert.Nil(t, err)

	err = writer.Close()
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/teacher/import", &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "917f6cd13b307b54c46e341b4c4e24ba"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetColleges(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/college/get", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "917f6cd13b307b54c46e341b4c4e24ba"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetMajorsByCollegeID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/college/4/get/major", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "917f6cd13b307b54c46e341b4c4e24ba"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetClassesByMajorID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/major/6/get/class", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "917f6cd13b307b54c46e341b4c4e24ba"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}
