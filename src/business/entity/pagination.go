package entity

type Pagination struct {
	CurrentPage   int64    `json:"current_page"` // for paginate in mysql
	TotalPages    int64    `json:"total_pages"` // for paginate in mysql
	TotalElements int64    `json:"total_elements"` // for paginate in mysql
	CursorStart   *string  `json:"cursor_start,omitempty"` // can use to pagination from mysql very fast performance
	CursorEnd     *string  `json:"cursor_end,omitempty"`
	HasNextPage   bool     `json:"has_next_page"`   // use to get other page from google APIS
	NextPageToken string   `json:"next_page_token"` // use to get other page from google APIS
}
