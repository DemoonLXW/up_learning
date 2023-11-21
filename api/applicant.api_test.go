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

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/stretchr/testify/assert"
)

func AddFilesToForm(writer *multipart.Writer, fnames []string, field string) error {
	for _, v := range fnames {
		f, err := os.Open(v)
		if err != nil {
			return err
		}
		defer f.Close()

		part, err := writer.CreateFormFile(field, filepath.Base(f.Name()))
		if err != nil {
			return err
		}
		_, err = io.Copy(part, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestAddProject(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	project := entity.ToAddProject{
		Title:               "api title2",
		Goal:                "api goal2",
		Principle:           "api principle2",
		ProcessAndMethod:    "api process and method 2",
		Step:                "api step2",
		ResultAndConclusion: "api result and conclusion 2",
		Requirement:         "api requirement2",
	}

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)

	fnames := []string{"./main.go", "./workflow.config.yaml"}

	err = AddFilesToForm(writer, fnames, "attachments")
	assert.Nil(t, err)

	writer.WriteField("title", project.Title)
	writer.WriteField("goal", project.Goal)
	writer.WriteField("principle", project.Principle)
	writer.WriteField("process_and_method", project.ProcessAndMethod)
	writer.WriteField("step", project.Step)
	writer.WriteField("result_and_conclusion", project.ResultAndConclusion)
	writer.WriteField("requirement", project.Requirement)

	recorder := httptest.NewRecorder()
	// body := bytes.NewReader([]byte(`{"title": "new project1", "goal": "new goal1", "requirement": "new requirement1"}`))
	// body := bytes.NewReader([]byte(`{}`))
	req, _ := http.NewRequest(http.MethodPost, "/applicant/project/add", &b)
	// req.Header.Add("Content-Type", binding.MIMEJSON)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	err = writer.Close()
	assert.Nil(t, err)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "f052dda02f37d4a5b5585e52e9719260"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestRemoveProjectsByIds(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"ids":[5]}`))
	req, _ := http.NewRequest(http.MethodPost, "/project/remove", body)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "004e9f57ec5bd6b100d211fd3e3b7d23"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestModifyAProject(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)

	project := entity.ToAddProject{
		Title:               "api update title4",
		Goal:                "api update goal4",
		Principle:           "api update principle4",
		ProcessAndMethod:    "api update process and method 4",
		Step:                "api update step4",
		ResultAndConclusion: "api update result and conclusion 4",
		Requirement:         "api update requirement4",
	}

	writer.WriteField("id", "4")

	writer.WriteField("title", project.Title)
	writer.WriteField("goal", project.Goal)
	writer.WriteField("principle", project.Principle)
	writer.WriteField("process_and_method", project.ProcessAndMethod)
	writer.WriteField("step", project.Step)
	writer.WriteField("result_and_conclusion", project.ResultAndConclusion)
	writer.WriteField("requirement", project.Requirement)
	writer.WriteField("delete_file_ids", "9")
	writer.WriteField("delete_file_ids", "12")

	fnames := []string{"./workflow.config.yaml", "./main.go"}
	err = AddFilesToForm(writer, fnames, "add_file")
	assert.Nil(t, err)

	body := bytes.NewReader([]byte(`{"id": 32, "title": "new project1", "goal": "new goal1", "requirement": "new requirement1"}`))

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/applicant/project/modify", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	err = writer.Close()
	assert.Nil(t, err)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "c07ecb8dd6d63d50ab834f0303c1e282"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetAProjectById(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/teacher/project/get/4", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "19392940f31df4e7af91269133cf330c"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetMyProjectList(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	query := "?current=1&page_size=5"
	req, _ := http.NewRequest(http.MethodGet, "/applicant/project/getlist"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "b827b0201fdc8a1bd14909d2598f3e52"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

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

func TestSubmitProjectForReview(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/applicant/project/submit-for-review?id=5", nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "4"}
	token_cookie := &http.Cookie{Name: "token", Value: "23ee603cc127e5018d899444f8d62b4e"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetReviewProjectRecordByProjectID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	query := "?current=1&page_size=1&project_id=5&sort=startTime&order=true"

	req, _ := http.NewRequest(http.MethodGet, "/applicant/review-project-record/get-list"+query, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "4"}
	token_cookie := &http.Cookie{Name: "token", Value: "61e5787bd6c0075729a3bbfcdbfffbcb"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}

func TestGetAReviewProjectRecordDetailByID(t *testing.T) {
	app, err := CreateTestApp()
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	id := "5bc739da-86c0-11ee-8812-0242aec148e6"

	req, _ := http.NewRequest(http.MethodGet, "/applicant/review-project-record/get/"+id, nil)
	uid_cookie := &http.Cookie{Name: "uid", Value: "4"}
	token_cookie := &http.Cookie{Name: "token", Value: "f8c68778885e2e8abba877c0eebfed79"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}
