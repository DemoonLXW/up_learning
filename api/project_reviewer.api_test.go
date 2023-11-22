package api_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTaskListOfPlatformReviewer(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?current=1&page_size=5"
	req, _ := http.NewRequest(http.MethodGet, "/project-reviewer/task/get-list"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "5"}
	token_cookie := &http.Cookie{Name: "token", Value: "333e0f34f6a704c8c54a5bcc665fd135"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetATaskDetailByID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	taskID := "5bdc6f99-86c0-11ee-8812-0242aec148e6"
	req, _ := http.NewRequest(http.MethodGet, "/project-reviewer/task/get/"+taskID, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "5"}
	token_cookie := &http.Cookie{Name: "token", Value: "333e0f34f6a704c8c54a5bcc665fd135"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestReviewProjectByTaskID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"id":"5bdc6f99-86c0-11ee-8812-0242aec148e6", "action": 1}`))
	req, _ := http.NewRequest(http.MethodPost, "/project-reviewer/task/complete", body)
	uid_cookie := &http.Cookie{Name: "uid", Value: "5"}
	token_cookie := &http.Cookie{Name: "token", Value: "3292c05cca124204d8003c9615c309d3"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}
