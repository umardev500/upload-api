package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/upload-api/domain"
)

func Ok(c *fiber.Ctx, status int, message string, data interface{}) error {
	payload := domain.SuccessResponse{
		Status:  status,
		Success: true,
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(payload)
}

func Failed(c *fiber.Ctx, status int, message string) error {
	payload := domain.FailedResponse{
		Status:  status,
		Success: false,
		Message: message,
	}
	return c.Status(status).JSON(payload)
}
