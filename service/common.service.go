package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"regexp"

	"github.com/DemoonLXW/up_learning/database/ent"
)

type CommonService struct {
	DB *ent.Client
}

func (serv *CommonService) SaveImportedFile(file *multipart.FileHeader, dir, prefix string) (*os.File, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("save imported file open file failed: %w", err)
	}
	defer src.Close()

	if err = os.MkdirAll(dir, 0750); err != nil {
		return nil, fmt.Errorf("save imported file mkdir failed: %w", err)
	}

	out, err := os.CreateTemp(dir, prefix)
	if err != nil {
		return nil, fmt.Errorf("save imported file create temp failed: %w", err)
	}

	_, err = io.Copy(out, src)
	if err != nil {
		return nil, fmt.Errorf("save imported file copy failed: %w", err)
	}
	defer out.Close()

	return out, nil
}

func (serv *CommonService) SaveUploadFile(file *multipart.FileHeader, dir string) (*os.File, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("save upload file open file failed: %w", err)
	}
	defer src.Close()

	re := regexp.MustCompile(`\.[^\.]+$`)
	res := re.FindStringSubmatch(file.Filename)
	suffix := res[0]

	if err = os.MkdirAll(dir, 0750); err != nil {
		return nil, fmt.Errorf("save upload file mkdir failed: %w", err)
	}

	h := md5.New()
	if _, err := io.Copy(h, src); err != nil {
		return nil, err
	}

	fname := fmt.Sprintf("%x", h.Sum(nil))

	out, err := os.OpenFile(dir+"/"+fname+suffix, os.O_RDWR|os.O_CREATE, 0750)
	if err != nil {
		return nil, fmt.Errorf("save upload file create out file failed: %w", err)
	}

	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("save upload file reset seek failed: %w", err)
	}

	_, err = io.Copy(out, src)
	if err != nil {
		return nil, fmt.Errorf("save upload file copy failed: %w", err)
	}
	defer out.Close()

	return out, nil
}
