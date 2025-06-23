package chunk

import (
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/lu4p/cat/docxtxt"
	"github.com/unidoc/unioffice/document"
)

type chunkedDoc struct {
	path   string
	file   *os.File
	text   string
	chunks []string
}

func ChunkDoc(path string) error {
	var doc chunkedDoc
	var err error
	doc.path = path
	// doc.file, err = os.Open(path)
	// if err != nil {
	// 	return fmt.Errorf("couldn't retrieve document at path %v, %v", path, err)
	// }
	err = doc.extractAllTextWithCatDoc()
	if err != nil {
		return fmt.Errorf("couldn't chunk document at path %v, %v", path, err)
	}
	fmt.Println("extracted text:", doc.text)
	return nil
}

func (doc *chunkedDoc) extractAllTextWithUniDoc() error {
	uniDoc, err := document.Open(doc.path)
	if err != nil {
		return fmt.Errorf("could not extract text from doc, %v", err)
	}
	defer uniDoc.Close()
	doc.text = uniDoc.ExtractText().Text()
	return nil
}

func (doc *chunkedDoc) extractAllTextWithCatDoc() error {
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
