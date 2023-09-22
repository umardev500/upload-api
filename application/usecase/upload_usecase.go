package usecase

import (
	"fmt"
	"mime/multipart"

	"github.com/umardev500/upload-api/domain"
)

type uploadUsecase struct{}

func NewUploadUsecase() domain.UploadUsecase {
	return &uploadUsecase{}
}

func (uc uploadUsecase) SaveToFileChunk(chunk *multipart.FileHeader, chunkpath string) error {
	f, err := chunk.Open()
	if err != nil {
		return err
	}

	fmt.Println(f)

	return nil
}
