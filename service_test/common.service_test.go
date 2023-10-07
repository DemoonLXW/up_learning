package service_test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func CreateTestCommonService() (*service.CommonService, error) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.CommonService)
	db, err := injection.ProvideDataBase()
	if err != nil {
		return nil, err
	}
	serv.DB = db

	return serv, nil
}

func TestSaveImportedFile(t *testing.T) {
	serv, err := CreateTestCommonService()
	assert.Nil(t, err)

	f, err := os.Open("../domain.config.json")
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

	req, err := http.NewRequest(http.MethodPost, "localhost:8080", &b)
	assert.Nil(t, err)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	_, fh, err := req.FormFile("upload")
	assert.Nil(t, err)

	dir := "../temp/importClass"
	prefix := "liu"

	o, err := serv.SaveImportedFile(fh, dir, prefix)
	assert.Nil(t, err)
	fmt.Println(o.Name())
	// os.Remove(o.Name())
}
