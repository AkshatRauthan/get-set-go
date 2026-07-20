package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ReadFilesAtOnce uses os.ReadFile to read the complete file and load its content into Ram as bytes buffer.
func ReadFilesAtOnce() ([]byte, error) {
	filePath := "./test-file.txt"

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("ReadFilesAtOnce-1: file %s does not exist", filePath)
		}
		return nil, fmt.Errorf("ReadFilesAtOnce-2: %v", err)
	}

	return data, nil
}

// ReadFilesManually is same as above function, but here we are doing everything manually.
// This is helpful when we need to embed custom logic in between opening and reading the file contents. (Like checking file size etc.)
func ReadFilesManually() ([]byte, error) {
	filePath := "./test-file.txt"

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("ReadFilesManually-1: Error: file %s does not exist", filePath)
		}
		return nil, fmt.Errorf("ReadFilesManually-2: Error: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFilesManually-3: %v", err)
	}

	return data, nil
}

// ReadFilesLineByLine reads the content of files line by line ensuring that we don't load the complete file
// directly into our memory at once.
func ReadFilesLineByLine() error {
	filePath := "./test-file.txt"

	f, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("ReadFilesLineByLine-1: file %s does not exist", filePath)
		}
		return fmt.Errorf("ReadFilesLineByLine-2: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Scan() stops on EOF (returns false) but doesn't treat it as an error.
	// Any other error (e.g. I/O failure) needs to be checked explicitly.
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
