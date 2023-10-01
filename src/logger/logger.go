package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init() {
	// Create or open a log file for writing
	// logFile, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	logrus.Fatalf("Failed to open log file: %v", err)
	// }

	// Set the log output to the file
	logrus.SetOutput(os.Stdout) // change os.Stdout to logFile to write to logfile.log

	// Set log level
	logrus.SetLevel(logrus.DebugLevel)

	// Optionally, you can also set the log formatter, e.g., JSONFormatter
	// logrus.SetFormatter(&logrus.JSONFormatter{})
}
