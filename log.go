// Package log is togo's configurable logging plugin: levels, text/JSON format,
// and file output. It overrides the kernel's baseline logger. Error trackers
// (Sentry, GlitchTip, …) ship as separate plugins that subscribe to the kernel
// "error" hook — this package only configures the slog sink.
//
// Install: `togo install togo-framework/log` (blank-import registers it).
package log

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("log", togo.PriorityCore, func(k *togo.Kernel) error {
		k.Log = New()
		return nil
	})
}

// New builds a slog logger from the environment:
//
//	LOG_LEVEL   debug|info|warn|error      (default info)
//	LOG_FORMAT  text|json                  (default text)
//	LOG_FILE    path                       (default stderr; appends)
//	LOG_SERVICE name                       (added as a "service" attribute)
func New() *slog.Logger {
	level := slog.LevelInfo
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	var w io.Writer = os.Stderr
	if path := os.Getenv("LOG_FILE"); path != "" {
		if f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644); err == nil {
			w = f
		}
	}

	opts := &slog.HandlerOptions{Level: level}
	var h slog.Handler = slog.NewTextHandler(w, opts)
	if strings.ToLower(os.Getenv("LOG_FORMAT")) == "json" {
		h = slog.NewJSONHandler(w, opts)
	}
	l := slog.New(h)
	if svc := os.Getenv("LOG_SERVICE"); svc != "" {
		l = l.With("service", svc)
	}
	return l
}
