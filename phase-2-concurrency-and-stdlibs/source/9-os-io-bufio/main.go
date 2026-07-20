package main

import (
	"fmt"
	"os"
)

func main() {

	// AtOnce Operations: standard library operations
	if err := WriteFileAtOnce(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	data, err := ReadFilesAtOnce()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		fmt.Printf("File Contents:\n%s\n", data)
	}

	// Manual Operations: manual operations for custom logic addition
	if err := WriteFileManually(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	data, err = ReadFilesManually()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		fmt.Printf("\n\nFile Contents:\n%s\n", data)
	}

	// Line By Line Operations: for buffered tasks
	if err := WriteFileLineByLine(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	fmt.Println("\n\nFile Contents:")
	if err = ReadFilesLineByLine(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
