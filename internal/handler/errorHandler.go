package handler

import (
	"fmt"
	"log"
	"os"
)

// ErrorHandler logs the error with a message and returns a wrapped error.
func ErrorHandler(err error, message string) error {
	errorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		errorLogger.Println(message, err)
		// Wrap the original error so you don't lose context
		return fmt.Errorf("%s: %w", message, err)
	}
	return fmt.Errorf(message)
}
