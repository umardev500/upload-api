package domain

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type UploadHandler interface {
	UploadChunk(c *fiber.Ctx) error
}

type UploadUsecase interface {
	SaveToFileChunk(file *multipart.FileHeader, chunkpath string) error
}
