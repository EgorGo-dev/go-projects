package main

import (
    "log/slog"
)

func LogHTTPRequest(logger *slog.Logger, method, path string, status int, durationMs int64) {
    logger.Info("http request",
        slog.String("method", method),
        slog.String("path", path),
        slog.Int("status", status),
        slog.Int64("duration_ms", durationMs),
    )
}