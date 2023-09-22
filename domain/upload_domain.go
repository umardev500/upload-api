package domain

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type UploadHandler interface {
	UploadChunk(c *fiber.Ctx) error
}

type UploadUsecase interface {
	// SaveToFile save file uploaded to storage
	//
	// Params:
	//   - file *multipart.FileHeader
	//   - filename string
	//
	// Return: error
	SaveToFile(file *multipart.FileHeader, chunkpath string) error
	// Reassemble reassemble the file to actually file
	//
	// Params:
	//   - root string - is root of file tree
	//   - dest string - is file destination output
	//
	// Return: error
	Reassemble(root, dest string) error
}
