package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"regexp"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/entity"
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

func (serv *CommonService) CreateFile(ctx context.Context, client *ent.Client, toCreates []*entity.ToAddFile) error {
	if ctx == nil || client == nil {
		return fmt.Errorf("context or client is nil")
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := client.File.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(file.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create file find a deleted file id query failed: %w", err)
			}
			break
		}

		num, err := client.File.Update().Where(file.And(
			file.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(file.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearModifiedTime().
			ClearDeletedTime().
			SetUID(toCreates[current].UID).
			SetName(toCreates[current].Name).
			SetPath(toCreates[current].Path).
			SetSize(toCreates[current].Size).
			SetIsDisabled(false).
			Save(ctx)

		if err != nil {
			return fmt.Errorf("create file: %w", err)
		}
		if num == 0 {
			return fmt.Errorf("create file update deleted file affect 0 row")
		}
	}
	if current < length {
		num, err := client.File.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create file count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.FileCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = client.File.Create().
				SetID(uint32(num + i + 1)).
				SetUID(toCreates[current+i].UID).
				SetName(toCreates[current+i].Name).
				SetPath(toCreates[current+i].Path).
				SetSize(toCreates[current+i].Size)
		}
		err = client.File.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return fmt.Errorf("create file failed: %w", err)
		}
	}

	return nil
}
