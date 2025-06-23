package cli

import (
	"fmt"
	"os"
	"strings"
)

// Args holds all CLI arguments
type Args struct {
	Filename string
}

// GetUserArgs parses and validates command line arguments
func GetUserArgs() (*Args, error) {
	// Parse command line arguments to get the file name
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("usage: go run main.go <filename.docx>")
	}

	filename := os.Args[1]

	// Validate that the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", filename)
	}

	// Validate that it is a .docx file type
	if !strings.HasSuffix(filename, ".docx") {
		return nil, fmt.Errorf("file must be a .docx file, got: %s", filename)
	}

	return &Args{
		Filename: filename,
	}, nil
}
