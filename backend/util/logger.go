package util

import (
	"log/slog"
	"os"
)

func init() {
	// Specify the log file path
	logFilePath := "app.log"

	// Open the log file
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	logger := slog.New(
		slog.NewJSONHandler(
			file,
			&slog.HandlerOptions{
				AddSource: true,
			}))
	slog.SetLogLoggerLevel(slog.LevelInfo)
	slog.SetDefault(logger)
	// Optionally, add common attributes to all log entries
	logger = logger.With(
		slog.String("service", "my_service"),
		slog.Int("version", 1),
	)
}

// Usage:
// util.Logger.Info("User login", slog.String("username", "johndoe"))
