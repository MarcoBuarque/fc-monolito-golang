package types

type Status string

type DocumentType string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Declined Status = "declined"

	RG   DocumentType = "RG"
	CPF  DocumentType = "CPF"
	CNPJ DocumentType = "CNPJ"
)

func (data Status) ToString() string {
	return string(data)
}

func (data DocumentType) ToString() string {
	return string(data)
}
