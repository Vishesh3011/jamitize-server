package config

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"syscall"
)

type loggerConfig struct {
	url string
}

func newLoggerConfig() (*loggerConfig, error) {
	url, found := syscall.Getenv("LOG_URL")
	if !found {
		return nil, fmt.Errorf("environment variable LOG_URL not set")
	}
	return &loggerConfig{
		url: url,
	}, nil
}

func (l *loggerConfig) URL() string {
	return l.url
}

func NewLogger(url string) (*slog.Logger, error) {
	var logger *slog.Logger
	var writer io.Writer
	if url != "" {
		var err error
		writer, err = net.Dial("udp", url)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to log server: %w", err)
		}
	} else {
		writer = os.Stdout
	}
	logger = slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{}))
	return logger, nil
}
