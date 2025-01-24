package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Log struct {
	// stdout *log.Logger
	// stderr *log.Logger
	newLog *zerolog.Logger
	flag   string
}

// Remove Soon
// func New(prefix string) *Log {
// 	// now := time.Now().Format("2006/01/02 15:04:05")
// 	return &Log{
// 		// stdout: log.New(os.Stdout, "["+now+"] [LOG]["+prefix+"]", log.Lshortfile),
// 		// stderr: log.New(os.Stderr, "["+now+"] [ERROR]["+prefix+"]", log.Lshortfile),
// 		stdout: log.New(os.Stdout, "[LOG]["+prefix+"]", log.Ldate|log.Ltime),
// 		stderr: log.New(os.Stderr, "[ERROR]["+prefix+"]", log.Ldate|log.Ltime),
// 	}
// }

// NewLogger creates and configures a new logger
func New(prefix string) *Log {

	zerolog.TimeFieldFormat = "2006/01/02 15:04:05"

	output := zerolog.ConsoleWriter{
		Out:           os.Stdout,
		TimeFormat:    "2006/01/02 15:04:05",
		PartsOrder:    []string{"time", "category", "message"},
		FieldsExclude: []string{"category"},
		FormatLevel: func(i any) string {
			return "[" + strings.ToUpper(i.(string)) + "]"
		},
		FormatTimestamp: func(i any) string {
			return "[" + i.(string) + "]"
		},
		FormatFieldName: func(i any) string {
			return ""
		},
		FormatFieldValue: func(i any) string {
			if category, ok := i.(string); ok && category != "" {
				return "[" + category + "]"
			}
			return ""
		},
		FormatMessage: func(i any) string {
			return i.(string)
		},
	}

	logger := zerolog.New(output).With().Timestamp().Logger()

	return &Log{
		newLog: &logger,
		flag:   prefix,
	}
}
