package usecase

import (
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

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

func (uc uploadUsecase) Reassemble(root, dest string) error {
	var output []byte
	var ext string

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// skip root
		if path == root {
			return nil
		}
		// get file format
		if ext == "" {
			ext = filepath.Ext(path)
		}

		// open the file
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		// read all content
		fileByte, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		// append content to output
		output = append(output, fileByte...)

		// remove temporary file
		os.Remove(path)

		return nil
	})

	// chek for directory
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		os.MkdirAll(dest, os.ModePerm)
	}

	now := time.Now().UTC().UnixNano()
	filename := filepath.Join(dest, fmt.Sprintf("%d%s", now, ext))
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = out.Write(output)
	if err != nil {
		return err
	}

	return nil
}
