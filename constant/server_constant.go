package constant

type ResponseStatus int
type Headers int
type General int

// Constant Api
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
)

var (
	responseStatuses = [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}
	responseMessages = [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}
)

func (r ResponseStatus) GetResponseStatus() string {
	return responseStatuses[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return responseMessages[r-1]
}
