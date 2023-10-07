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
