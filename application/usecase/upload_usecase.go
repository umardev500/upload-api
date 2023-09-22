package usecase

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/umardev500/upload-api/domain"
)

type uploadUsecase struct{}

func NewUploadUsecase() domain.UploadUsecase {
	return &uploadUsecase{}
}

func (uc uploadUsecase) SaveToFile(file *multipart.FileHeader, filename string) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, f)
	if err != nil {
		return err
	}

	return nil
}
