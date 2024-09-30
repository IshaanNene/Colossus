// logger.go
package pkg

import (
	"fmt"
	"os"
	"sync"
)

// Logger represents a simple logging utility
type Logger struct {
	mu   sync.Mutex
	file *os.File
}

// NewLogger creates a new Logger instance
func NewLogger(filePath string) (*Logger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return &Logger{
		file: file,
	}, nil
}

// Log writes a log entry to the log file
func (l *Logger) Log(message string) {
	l.mu.Lock() // Acquire lock for thread-safe logging
	defer l.mu.Unlock()
	_, err := l.file.WriteString(fmt.Sprintf("%s\n", message))
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

// Close closes the log file
func (l *Logger) Close() {
	l.file.Close()
}
