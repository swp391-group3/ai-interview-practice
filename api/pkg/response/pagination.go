package response

// Page is a reusable offset-paginated response payload.
type Page[T any] struct {
	Items      []T        `json:"items"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Total  int64 `json:"total"`
}
