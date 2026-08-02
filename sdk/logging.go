package sdk

import (
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"strings"

	"github.com/lmittmann/tint"
)

// HumanLogHandlerOptions configures NewHumanLogHandler.
type HumanLogHandlerOptions struct {
	Level   slog.Leveler
	AddTime bool
	NoColor bool
}

// NewDebugLogger returns a human-readable debug logger for SDK examples and CLIs.
func NewDebugLogger(w io.Writer) *slog.Logger {
	return slog.New(NewHumanLogHandler(w, &HumanLogHandlerOptions{Level: slog.LevelDebug}))
}

// NewHumanLogHandler returns a slog handler that prints SDK events for humans.
func NewHumanLogHandler(w io.Writer, opts *HumanLogHandlerOptions) slog.Handler {
	if w == nil {
		w = io.Discard
	}
	level := slog.Leveler(slog.LevelInfo)
	addTime := false
	noColor := false
	if opts != nil {
		if opts.Level != nil {
			level = opts.Level
		}
		addTime = opts.AddTime
		noColor = opts.NoColor
	}
	return tint.NewHandler(w, &tint.Options{
		Level:       level,
		TimeFormat:  "15:04:05",
		NoColor:     noColor,
		ReplaceAttr: humanLogReplaceAttr(addTime, noColor),
	})
}

func humanLogReplaceAttr(addTime bool, noColor bool) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, attr slog.Attr) slog.Attr {
		if attr.Key == slog.TimeKey && len(groups) == 0 && !addTime {
			return slog.Attr{}
		}

		attr.Value = attr.Value.Resolve()
		if attr.Value.Kind() == slog.KindString && attr.Value.String() == "" && attr.Key != slog.MessageKey {
			return slog.Attr{}
		}
		if noColor {
			return attr
		}
		if len(groups) == 0 {
			switch attr.Key {
			case slog.LevelKey:
				if level, ok := attr.Value.Any().(slog.Level); ok {
					return slog.String(slog.LevelKey, styledHumanLogLevel(level))
				}
			case slog.MessageKey:
				if msg := attr.Value.String(); msg != "" {
					return slog.String(slog.MessageKey, boldHumanLogText(msg))
				}
			}
		}
		if isHumanLogErrorKey(attr.Key) {
			return tint.Attr(9, slog.String(attr.Key, boldHumanLogText(formatHumanLogAttrValue(attr.Value))))
		}
		if color, ok := humanLogAttrColor(attr.Key); ok {
			return tint.Attr(color, attr)
		}
		return attr
	}
}

func styledHumanLogLevel(level slog.Level) string {
	label := humanLogLevelLabel(level)
	switch {
	case level < slog.LevelInfo:
		return ansiHumanLogBoldCyan + label + ansiHumanLogReset
	case level < slog.LevelWarn:
		return ansiHumanLogBoldGreen + label + ansiHumanLogReset
	case level < slog.LevelError:
		return ansiHumanLogBoldYellow + label + ansiHumanLogReset
	default:
		return ansiHumanLogBoldRed + label + ansiHumanLogReset
	}
}

func humanLogLevelLabel(level slog.Level) string {
	switch {
	case level < slog.LevelInfo:
		return humanLogLevelLabelWithDelta("DBG", level-slog.LevelDebug)
	case level < slog.LevelWarn:
		return humanLogLevelLabelWithDelta("INF", level-slog.LevelInfo)
	case level < slog.LevelError:
		return humanLogLevelLabelWithDelta("WRN", level-slog.LevelWarn)
	default:
		return humanLogLevelLabelWithDelta("ERR", level-slog.LevelError)
	}
}

func humanLogLevelLabelWithDelta(base string, delta slog.Level) string {
	switch {
	case delta > 0:
		return fmt.Sprintf("%s+%d", base, delta)
	case delta < 0:
		return fmt.Sprintf("%s%d", base, delta)
	default:
		return base
	}
}

func humanLogAttrColor(key string) (uint8, bool) {
	switch {
	case key == "operation":
		return 13, true
	case strings.HasSuffix(key, "_id") || strings.HasSuffix(key, ".id"):
		return 12, true
	case strings.Contains(key, "url") || strings.Contains(key, "host") || strings.Contains(key, "path"):
		return 14, true
	case strings.Contains(key, "phase") || strings.Contains(key, "status") || strings.Contains(key, "state"):
		return 11, true
	case strings.Contains(key, "duration") || strings.Contains(key, "count") || strings.Contains(key, "size"):
		return 10, true
	default:
		return 0, false
	}
}

func isHumanLogErrorKey(key string) bool {
	key = strings.ToLower(key)
	return key == "err" || key == "error" || strings.HasSuffix(key, ".err") || strings.HasSuffix(key, ".error")
}

func formatHumanLogAttrValue(value slog.Value) string {
	switch value.Kind() {
	case slog.KindString:
		return value.String()
	default:
		return fmt.Sprint(value.Any())
	}
}

func boldHumanLogText(text string) string {
	return ansiHumanLogBold + text + ansiHumanLogBoldReset
}

const (
	ansiHumanLogBold       = "\x1b[1m"
	ansiHumanLogBoldCyan   = "\x1b[1;96m"
	ansiHumanLogBoldGreen  = "\x1b[1;92m"
	ansiHumanLogBoldYellow = "\x1b[1;93m"
	ansiHumanLogBoldRed    = "\x1b[1;91m"
	ansiHumanLogBoldReset  = "\x1b[22m"
	ansiHumanLogReset      = "\x1b[0m"
)

func safeURLForLog(rawURL string) string {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return ""
	}

	parsed, err := url.Parse(rawURL)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return rawURL
	}

	parsed.User = nil
	parsed.RawQuery = ""
	parsed.Fragment = ""
	return parsed.String()
}
