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
	Duplicated
	ValidateError
	BadRequest
	RequiredQuery
	DataIsExit
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED", "DUPLICATED", "ValidateError", "BAD_REQUEST",
		"REQUIRED_QUERY", "DATA_IS_EXIT"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized", "Duplicated", "ValidateError", "BadRequest", "RequiredQuery", "DataIsExit"}[r-1]
}
