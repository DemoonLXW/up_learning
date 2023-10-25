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

	"github.com/stretchr/testify/assert"
)

func TestUploadDocumentImage(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	f, err := os.Open("../import.xlsx")
	assert.Nil(t, err)
	defer f.Close()

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("upload", filepath.Base(f.Name()))
	assert.Nil(t, err)

	_, err = io.Copy(part, f)
	assert.Nil(t, err)
	err = writer.Close()
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/file/uploadimage", &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "ef1416f4191e70f72edc98782effff4b"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestDownloadFile(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?ids=12&ids=9"

	req, _ := http.NewRequest(http.MethodGet, "/common/file/download"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "9ab6785cf9bae4c5a020bf603c4d9a5f"}
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

	down, err := os.Create("temp/test/" + res[1])
	assert.Nil(t, err)

	_, err = recorder.Body.WriteTo(down)
	assert.Nil(t, err)
}
