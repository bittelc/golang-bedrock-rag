package chunk

import (
	"fmt"
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/lu4p/cat/docxtxt"

	// "github.com/unidoc/unioffice/document"
	"github.com/srinathh/gooxml/document"
)

type chunkDoc struct {
	path   string
	file   *os.File
	text   string
	chunks []string
}

func ChunkDoc(path string) error {
	var doc chunkDoc
	var err error
	doc.path = path
	err = doc.extractAllTextWithForekdUniDoc()
	if err != nil {
		return fmt.Errorf("couldn't chunk document at path %v, %v", path, err)
	}
	fmt.Println("extracted text:", doc.text)
	return nil
}

// func (doc *chunkDoc) extractAllTextWithUniDoc() error {
// 	uniDoc, err := document.Open(doc.path)
// 	if err != nil {
// 		return fmt.Errorf("could not extract text from doc, %v", err)
// 	}
// 	defer uniDoc.Close()
// 	doc.text = uniDoc.ExtractText().Text()
// 	return nil
// }

func (d *chunkDoc) extractAllTextWithForekdUniDoc() error {
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

func (doc *chunkDoc) extractAllTextWithCatDoc() error {
	content, err := os.ReadFile(doc.path)
	if err != nil {
		return fmt.Errorf("could not detect mime type, %v", err)
	}
	mime := mimetype.Detect(content)
	fmt.Println("cat will process this doc as mimetype:", mime.String())
	txt, err := docxtxt.ToStr(doc.path)
	if err != nil {
		return fmt.Errorf("could not extract text from doc, %v", err)
	}
	doc.text = txt
	return nil
}

func (doc *chunkDoc) extractAllTextWithCatTryDifferentWay() error {
	f, err := os.Open(doc.path)
	if err != nil {
		return err
	}
	defer f.Close()
	doc.text, err = docxtxt.ToStr(doc.path)
	if err != nil {
		return err
	}
	return nil
}
