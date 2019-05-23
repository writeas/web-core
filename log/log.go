// Package log prints different kinds of log messages in the formats we like.
package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {
	InfoLog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

// Info logs an informational message to Stdout.
func Info(s string, v ...interface{}) {
	InfoLog.Printf(s, v...)
}

// Error logs an error to Stderr.
func Error(s string, v ...interface{}) {
	// Include original caller information
	_, file, line, _ := runtime.Caller(1)

	// Determine short filename (from standard log package)
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short

	ErrorLog.Printf(fmt.Sprintf("%s:%d: ", short, line)+s, v...)
}
