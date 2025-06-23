package chunk

import (
	"fmt"
	"os"
	"strings"

	// "github.com/unidoc/unioffice/document"
	"github.com/srinathh/gooxml/document"
)

const embeddingVectorSize = 500

type chunkDoc struct {
	path   string
	file   *os.File
	text   string
	chunks []string
}

func ChunkDoc(path string) (*chunkDoc, error) {
	var doc chunkDoc
	doc.path = path
	err := doc.extractAllText()
	if err != nil {
		return &doc, fmt.Errorf("couldn't chunk document at path %v, %v", path, err)
	}

	err = doc.chunkText()

	return &doc, nil
}

func (d *chunkDoc) extractAllText() error {
	// Open the file
	file, err := os.Open(d.path)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("could not get file info: %v", err)
	}

	// Read the document using gooxml
	doc, err := document.Read(file, fileInfo.Size())
	if err != nil {
		return fmt.Errorf("could not read document: %v", err)
	}

	// Extract text from all paragraphs
	var textBuilder strings.Builder
	for _, par := range doc.Paragraphs() {
		for _, run := range par.Runs() {
			textBuilder.WriteString(run.Text())
		}
		// Add newline after each paragraph
		textBuilder.WriteString("\n")
	}

	// Store the extracted text
	d.text = textBuilder.String()
	return nil
}

func (d *chunkDoc) chunkText() error {
	words := strings.Fields(d.text)
	var chunks []string
	for i := 0; i < len(words); i += embeddingVectorSize {
		end := i + embeddingVectorSize
		if end > len(words) {
			end = len(words)
		}
		chunks = append(chunks, strings.Join(words[i:end], " "))
	}
	d.chunks = chunks
	return nil
}
