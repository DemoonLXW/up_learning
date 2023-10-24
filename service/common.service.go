package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
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

func (serv *CommonService) SaveUploadFile(file *multipart.FileHeader, dir string) (*entity.ToAddFile, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("save upload file open file failed: %w", err)
	}
	defer src.Close()

	re := regexp.MustCompile(`\.[^\.]+$`)
	res := re.FindStringSubmatch(file.Filename)
	suffix := res[0]

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("save upload file get abs directory failed: %w", err)
	}

	if err = os.MkdirAll(absDir, 0750); err != nil {
		return nil, fmt.Errorf("save upload file mkdir failed: %w", err)
	}

	h := md5.New()
	if _, err := io.Copy(h, src); err != nil {
		return nil, err
	}

	fname := fmt.Sprintf("%x", h.Sum(nil))

	out, err := os.OpenFile(absDir+"/"+fname+suffix, os.O_RDWR|os.O_CREATE, 0750)
	if err != nil {
		return nil, fmt.Errorf("save upload file create out file failed: %w", err)
	}
	defer out.Close()

	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("save upload file reset seek failed: %w", err)
	}

	_, err = io.Copy(out, src)
	if err != nil {
		return nil, fmt.Errorf("save upload file copy failed: %w", err)
	}

	info, err := out.Stat()
	if err != nil {
		return nil, fmt.Errorf("save upload file get info failed: %w", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("save upload file get working directory failed: %w", err)
	}
	path, err := filepath.Rel(wd, out.Name())
	if err != nil {
		return nil, fmt.Errorf("save upload file get addfile path failed: %w", err)
	}

	var addFile entity.ToAddFile
	addFile.Name = file.Filename
	addFile.Path = path
	addFile.Size = info.Size()
	return &addFile, nil
}

func (serv *CommonService) CreateFile(ctx context.Context, client *ent.Client, toCreates []*entity.ToAddFile) ([]uint32, error) {
	if ctx == nil || client == nil {
		return nil, fmt.Errorf("context or client is nil")
	}

	current := 0
	length := len(toCreates)

	createIDs := make([]uint32, 0, length)

	for ; current < length; current++ {
		id, err := client.File.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(file.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, fmt.Errorf("create file find a deleted file id query failed: %w", err)
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
			return nil, fmt.Errorf("create file: %w", err)
		}
		if num == 0 {
			return nil, fmt.Errorf("create file update deleted file affect 0 row")
		}

		createIDs = append(createIDs, id)
	}
	if current < length {
		num, err := client.File.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return nil, fmt.Errorf("create file count failed: %w", err)
		}

		bulkLength := length - current
		// bulk := make([]*ent.FileCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			err = client.File.Create().
				SetID(uint32(num + i + 1)).
				SetUID(toCreates[current+i].UID).
				SetName(toCreates[current+i].Name).
				SetPath(toCreates[current+i].Path).
				SetSize(toCreates[current+i].Size).
				Exec(ctx)
			if err != nil {
				return nil, fmt.Errorf("create file failed: %w", err)
			}

			createIDs = append(createIDs, uint32(num+i+1))

		}
		// err = client.File.CreateBulk(bulk...).Exec(ctx)
	}

	return createIDs, nil
}

func (serv *CommonService) DeleteFile(ctx context.Context, client *ent.Client, toDeletes []*ent.File) error {
	if ctx == nil || client == nil {
		return fmt.Errorf("context or client is nil")
	}

	for i := range toDeletes {
		num, err := client.File.Update().
			Where(file.And(
				file.IDEQ(toDeletes[i].ID),
				func(s *sql.Selector) {
					s.Where(sql.IsNull(file.FieldDeletedTime))
				},
			)).
			ClearModifiedTime().
			SetDeletedTime(time.Now()).
			SetName("*" + toDeletes[i].Name).
			Save(ctx)

		if err != nil {
			return fmt.Errorf("delete file failed: %w", err)
		}
		if num == 0 {
			return fmt.Errorf("delete file affect 0 row: id[%d]", toDeletes[i].ID)
		}
	}

	return nil
}

func (serv *CommonService) FindFileByIDs(ctx context.Context, client *ent.Client, IDs []uint32) ([]*ent.File, error) {
	if ctx == nil || client == nil {
		return nil, fmt.Errorf("context or client is nil")
	}

	res, err := client.File.Query().Where(file.And(
		file.IDIn(IDs...),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(file.FieldDeletedTime))
		},
	)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("find file by ids failed: %w", err)
	}
	if len(res) != len(IDs) {
		return nil, fmt.Errorf("find file by ids failed: number of result is not enough")
	}

	return res, nil
}
