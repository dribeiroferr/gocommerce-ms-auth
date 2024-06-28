package commons

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// NewLogger sets up the logger with the desired settings
func NewLogger() *logrus.Logger {
	logger := logrus.New()

	// Set the output to stdout
	logger.Out = os.Stdout

	// Set the log level
	logger.SetLevel(logrus.DebugLevel)

	// Set the formatter to include colors based on log level
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339, // Custom timestamp format
	})

	return logger
}

// Log is the global logger instance
var Log = NewLogger()
