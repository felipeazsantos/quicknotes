package main

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func formatSource(source string) (string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	baseDir := filepath.Base(workDir)
	source = strings.ReplaceAll(source, "}", "")
	source = strings.ReplaceAll(source, " ", ":")
	return source[strings.Index(source, baseDir):], nil
}

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == "time" {
		value := time.Now().Format("2006-01-02T15:04:05")
		return slog.Attr{Key: a.Key, Value: slog.StringValue(value)}
	} else if a.Key == "source" {
		value, err := formatSource(a.Value.String())
		if err != nil {
			return a
		}
		return slog.Attr{Key: a.Key, Value: slog.StringValue(value)}
	}
	return a
}

func newLogger(out io.Writer, minLevel slog.Level) *slog.Logger {
	h := slog.NewTextHandler(out, &slog.HandlerOptions{
		AddSource:   true,
		Level:       minLevel,
		ReplaceAttr: replaceAttr,
	})

	return slog.New(h)
}
