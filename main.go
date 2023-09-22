package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/upload-api/application"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	app := application.Application{}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
