package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/upload-api/application/handler"
	"github.com/umardev500/upload-api/application/usecase"
)

func LoadRoutes(app *fiber.App) {
	app.Route("/api", loadApiRoutes)
}

func loadApiRoutes(router fiber.Router) {
	router.Route("/", loadUploadRoutes)
}

func loadUploadRoutes(router fiber.Router) {
	uuc := usecase.NewUploadUsecase()
	handler := handler.NewUploadHandler(uuc)
	router.Post("/:chunk_id/:chunk_total/:chunk_index", handler.UploadChunk)
}
