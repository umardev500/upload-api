package application

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/upload-api/routes"
)

type Application struct{}

func (a *Application) Start(ctx context.Context) error {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	routes.LoadRoutes(app)

	ch := make(chan error, 1)
	go func() {
		port := os.Getenv("PORT")
		err := app.Listen(port)
		if err != nil {
			ch <- err
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		app.ShutdownWithTimeout(10 * time.Second)
	}

	return nil
}
