// Package log prints different kinds of log messages in the formats we like.
package log

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {
	InfoLog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an informational message to Stdout.
func Info(s string, v ...interface{}) {
	InfoLog.Printf(s, v...)
}

// Error logs an error to Stderr.
func Error(s string, v ...interface{}) {
	ErrorLog.Printf(s, v...)
}
