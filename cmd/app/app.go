package app

import (
	"delivery/config"
	"delivery/internal/server"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) error {
	//Server
	server := server.NewServer(cfg)
	err := server.Start()
	if err != nil {
		slog.Error("can't start a server %w", err)
		return err
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("app - Run - signal: " + s.String())
	case err = <-server.Notify():
		slog.Error("app - Run - httpServer.Notify: %w", err)
	}

	// Shutdown
	err = server.Shutdown()
	if err != nil {
		slog.Error("app - Run - httpServer.Shoutdown: %w", err)
	}
	return nil
}
