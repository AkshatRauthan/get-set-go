package main

import (
	"bufio"
	"fmt"
	"os"
)

// HERE, in end of all operations I have used O_TRUNC => So all operations will truncate then write
// If we want to append simply replace it with O_APPEND

// WriteFileAtOnce uses os.WriteFile to write the complete content to a file in one call.
// It creates the file if it doesn't exist, or truncates it if it does.
func WriteFileAtOnce() error {
	str := "This is the result of WriteFileAtOnce operation"
	data := []byte(str)

	filePath := "./test-file.txt"
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("WriteFileAtOnce-1: %v", err)
	}
	return nil
}

// WriteFileManually is same as above function, but here we are doing everything manually.
// This is helpful when we need custom control over file creation flags (append, exclusive create, etc.)
// or want to do something between opening and writing (like locking, checking free disk space, etc.)
func WriteFileManually() error {
	str := "This is the result of WriteFileManually operation"
	data := []byte(str)

	filePath := "./test-file.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("WriteFileManually-1: Error: %v", err)
	}
	defer file.Close()

	n, err := file.Write(data)
	if err != nil {
		return fmt.Errorf("WriteFileManually-2: %v", err)
	}
	if n < len(data) {
		return fmt.Errorf("WriteFileManually-3: short write, wrote %d of %d bytes", n, len(data))
	}
	return nil
}

// WriteFileLineByLine writes content line by line using a buffered writer,
// which avoids allocating one large byte slice for the entire content upfront
// and flushes to disk efficiently rather than issuing a syscall per line.
func WriteFileLineByLine() error {
	lines := []string{
		"This is the result of WriteFileLineByLine operation Line1",
		"This is the result of WriteFileLineByLine operation Line2",
		"This is the result of WriteFileLineByLine operation Line3",
		"This is the result of WriteFileLineByLine operation Line4",
		"This is the result of WriteFileLineByLine operation Line5",
	}

	filePath := "./test-file.txt"
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("WriteFileLineByLine-1: Error: %v", err)
	}
	defer f.Close()

	// Appending current line input + end-line character
	writer := bufio.NewWriter(f)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("WriteFileLineByLine-2: %v", err)
		}
	}

	// Flush() is essential — bufio.Writer buffers in memory and only
	// writes to the underlying file when the buffer fills up or Flush is called.
	// Skipping this can silently lose the last batch of buffered data.
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("WriteFileLineByLine-3: %v", err)
	}
	return nil
}
