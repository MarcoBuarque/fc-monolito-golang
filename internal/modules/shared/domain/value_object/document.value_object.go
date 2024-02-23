package valueobject

import (
	"fmt"

	"github.com/MarcoBuarque/monolito/internal/modules/shared/types"
)

type Document struct {
	number       string
	documentType types.DocumentType
}

func NewDocument(number string, docType types.DocumentType) (Document, error) {
	if number == "" {
		return Document{}, fmt.Errorf("document: number cannot be empty")
	}

	if docType == "" {
		return Document{}, fmt.Errorf("document: document type cannot be empty")
	}

	return Document{number, docType}, nil
}

func (data Document) Number() string                   { return data.number }
func (data Document) DocumentType() types.DocumentType { return data.documentType }
