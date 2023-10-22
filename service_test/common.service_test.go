package service_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/entity"
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

func TestSaveUploadFile(t *testing.T) {
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

	dir := "../temp/uploadImage"

	o, err := serv.SaveUploadFile(fh, dir)
	assert.Nil(t, err)
	fmt.Println(o.Name())
	o, err = os.Open(o.Name())
	assert.Nil(t, err)

	info, err := o.Stat()
	assert.Nil(t, err)
	fmt.Println(info.Size())
}

func TestCreateFileByProjectID(t *testing.T) {
	serv, err := CreateTestCommonService()
	assert.Nil(t, err)

	projectID := uint32(2)

	file1 := entity.ToAddFile{
		UID:  1,
		Name: "test3.file",
		Path: "file/uid/test3.file",
		Size: 33,
	}

	file2 := entity.ToAddFile{
		UID:  1,
		Name: "test3.file",
		Path: "file/uid/test3.file",
		Size: 32,
	}

	adds := []*entity.ToAddFile{&file2, &file1}

	ctx := context.Background()
	client := serv.DB

	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return serv.CreateFileByProjectID(ctx, tx.Client(), adds, projectID)
	})
	// err = serv.CreateProject(nil, nil, adds)
	assert.Nil(t, err)
}

func TestDeleteFile(t *testing.T) {
	serv, err := CreateTestCommonService()
	assert.Nil(t, err)

	ctx := context.Background()
	client := serv.DB

	files := []*ent.File{
		{
			ID:   7,
			Name: "test2.file",
		},
		{
			ID:   9,
			Name: "test2.file",
		},
	}

	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return serv.DeleteFile(ctx, tx.Client(), files)
	})

	assert.Nil(t, err)

}

func TestFindFileByIDs(t *testing.T) {
	serv, err := CreateTestCommonService()
	assert.Nil(t, err)

	ctx := context.Background()
	client := serv.DB

	ids := []uint32{2, 8}

	files, err := serv.FindFileByIDs(ctx, client, ids)
	assert.Nil(t, err)
	for i := range files {
		fmt.Println(files[i])
	}

}
