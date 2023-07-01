package jsonlog

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// define a "level" type to represent the severity for the log entries
type Level int8

// the use of iota is controversial in production.
// The following constants will not be serialized
// (saved to a file and read back in), and they will
// not be export. This minimizes the risk for data corruption.
const (
	LevelInfo Level = iota // Has the value of 0
	LevelError
	LevelFatal
	LevelOff
)

// helper function to return human-readable string for severity level
func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

// This will be our custom logger type
type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

// function to return a new logger instance
func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

// The following helper methods are what will be used throughout the application
// to write logs at different levels.
func (l *Logger) PrintInfo(message string, properties map[string]string) {
	l.print(LevelInfo, message, properties)
}

func (l *Logger) PrintError(err error, properties map[string]string) {
	l.print(LevelError, err.Error(), properties)
}

func (l *Logger) PrintFatal(err error, properties map[string]string) {
	l.print(LevelFatal, err.Error(), properties)
	os.Exit(1)
}

func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	// Check base case, make sure that the security level is not below our minimum value
	if level < l.minLevel {
		return 0, nil
	}

	// Create a struct to hold the data for our log entry
	payload := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"string"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}

	// we want a stack trace for logs at the error and fatal levels
	if level >= LevelError {
		payload.Trace = string(debug.Stack())
	}

	var output []byte

	// we need to marshal our struct to JSON
	output, err := json.Marshal(payload)
	if err != nil {
		output = []byte(LevelError.String() + ": unable to marshal log message:" + err.Error())
	}

	// important to lock the mutex here. We do not want to writes to the output
	// destination concurrently. The result could be that the text from two
	// *different* log entries get mingled together.
	l.mu.Lock()
	defer l.mu.Unlock()

	// FINALLY, write the log entry
	return l.out.Write(append(output, '\n')) // new line added for readability

}

// helper method for our Logger type that satisfies the io.Writer inteface.
func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, string(message), nil)
}
