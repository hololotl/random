package internal

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	*Handler
	Host     string
	Port     string
	logger   *slog.Logger
	server   *http.Server
	stopChan chan os.Signal
}

func InitHttpServer(logger *slog.Logger, Host string, Port string, rtp float64) (*HttpServer, error) {
	logger.Info("InitHttpServer")
	handler := NewHandler(rtp, logger)
	return &HttpServer{
		Handler:  handler,
		Host:     Host,
		Port:     Port,
		logger:   logger,
		stopChan: make(chan os.Signal, 1),
	}, nil
}

func (hs *HttpServer) Start() {
	router := hs.registerHandlers()

	hs.server = &http.Server{
		Addr:    hs.Host + ":" + hs.Port,
		Handler: router,
	}

	signal.Notify(hs.stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		hs.logger.Info("HTTP server is starting", "addr", hs.server.Addr)
		if err := hs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			hs.logger.Error("ListenAndServe error", "err", err)
		}
	}()

	<-hs.stopChan
	hs.logger.Info("Shutdown signal received")
	hs.Stop()
}

func (hs *HttpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hs.logger.Info("Shutting down HTTP server...")
	if err := hs.server.Shutdown(ctx); err != nil {
		hs.logger.Error("HTTP server Shutdown", "err", err)
	} else {
		hs.logger.Info("HTTP server exited properly")
	}
}

func (hs *HttpServer) registerHandlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/get", hs.GetRandom).Methods("GET")

	hs.logger.Info("Routes registered", "host", hs.Host, "port", hs.Port)
	return router
}
