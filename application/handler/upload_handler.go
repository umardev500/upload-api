package handler

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/upload-api/domain"
)

type uploadHandler struct {
	usecase domain.UploadUsecase
}

func NewUploadHandler(usecase domain.UploadUsecase) domain.UploadHandler {
	return &uploadHandler{
		usecase: usecase,
	}
}

func (u *uploadHandler) UploadChunk(c *fiber.Ctx) error {
	chunk, err := c.FormFile("chunk")
	if err != nil {
		return Failed(c, fiber.StatusBadRequest, err.Error())
	}

	// get passed params
	chunkID := c.Params("chunk_id")
	chunkTotal, _ := strconv.Atoi(c.Params("chunk_total"))
	chunkIndex, _ := strconv.Atoi(c.Params("chunk_index"))

	tempDir := filepath.Join("./storage/temp", chunkID)
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		os.MkdirAll(tempDir, os.ModePerm)
	}
	chunkpath := filepath.Join(tempDir, chunk.Filename)

	err = u.usecase.SaveToFile(chunk, chunkpath)
	if err != nil {
		return Failed(c, fiber.StatusInternalServerError, err.Error())
	}

	// validate chunk index and chunk total
	// to check for last chunk is uploaded
	if chunkTotal == (chunkIndex + 1) {
		dest := filepath.Join("storage", "video")
		err = u.usecase.Reassemble(tempDir, dest)
		if err != nil {
			return Failed(c, fiber.StatusInternalServerError, err.Error())
		}
		return Ok(c, fiber.StatusCreated, "video reassembled successfuly", nil)
	}

	return Ok(c, fiber.StatusOK, "Upload success", nil)
}
