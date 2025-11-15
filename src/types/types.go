package tavo

// PaginationInfo represents pagination information
type PaginationInfo struct {
	Page    int  `json:"page"`
	Limit   int  `json:"limit"`
	Total   int  `json:"total"`
	Pages   int  `json:"pages"`
	HasNext bool `json:"has_next"`
	HasPrev bool `json:"has_prev"`
}

// PaginatedResponse represents a paginated response
type PaginatedResponse[T any] struct {
	Data       []T             `json:"data"`
	Pagination PaginationInfo `json:"pagination"`
}

// ListResponse represents a list response
type ListResponse[T any] struct {
	Data  []T `json:"data"`
	Count int `json:"count"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}
