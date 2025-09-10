package main

import (
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"test/internal"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log := setupLogger(envLocal)
	host := "localhost"
	port := "64333"
	rtp := flag.Float64("rtp", -1.0, "значение RTP обязательно")
	flag.Parse()
	if *rtp == -1.0 {
		log.Error("неверный формат запуска, используй go run . -rtp={значение} ")
		return
	}
	fmt.Println("dfd", *rtp)
	server, err := internal.InitHttpServer(log, host, port, *rtp)
	if err != nil {
		log.Error("error creating http server", "error", err)
		return
	}
	server.Start()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
