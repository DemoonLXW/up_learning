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
		UID:                 3,
		Title:               "api title1",
		Goal:                "api goal1",
		Principle:           "api principle1",
		ProcessAndMethod:    "api process and method 1",
		Step:                "api step1",
		ResultAndConclusion: "api result and conclusion 1",
		Requirement:         "api requirement1",
	}

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)

	fnames := []string{"./workflow.config.yaml", "./domain.config.json"}

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
	req, _ := http.NewRequest(http.MethodPost, "/project/add", &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	err = writer.Close()
	assert.Nil(t, err)
	uid_cookie := &http.Cookie{Name: "uid", Value: "1"}
	token_cookie := &http.Cookie{Name: "token", Value: "230eeb1aeff1d068a164183820992a71"}
	req.AddCookie(uid_cookie)
	req.AddCookie(token_cookie)
	app.ServeHTTP(recorder, req)

	resp := recorder.Result()
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(recorder.Body.String())
}
