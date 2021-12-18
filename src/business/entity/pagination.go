package entity

type Pagination struct {
	CurrentPage     int64    `json:"current_page"`
	TotalPages      int64    `json:"total_pages"`
	TotalElements   int64    `json:"total_elements"`
	SortBy          []string `json:"sort_by"`
	CursorStart     *string  `json:"cursor_start,omitempty"` // TODO improve pagination using cursor
	CursorEnd       *string  `json:"cursor_end,omitempty"`
}
