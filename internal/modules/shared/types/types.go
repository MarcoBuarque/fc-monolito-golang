package types

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Declined Status = "declined"
)

func (data Status) ToString() string {
	return string(data)
}
