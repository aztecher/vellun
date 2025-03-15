package logging

import (
	"log/slog"
	"os"
)

var slogHandlerOpts = &slog.HandlerOptions{
	AddSource: false,
	Level:     slog.LevelInfo,
}

var DefaultSlogLogger *slog.Logger = slog.New(
	slog.NewTextHandler(
		os.Stderr,
		slogHandlerOpts,
	),
)
