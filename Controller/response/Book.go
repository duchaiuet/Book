package response

import "Book/Database/Models"

type CreateBookResponse struct {
	Meta
	Data *Models.Book `json:"data"`
}
type ListBooksResponse struct {
	Meta
	Data []*Models.Book `json:"data"`
	Pagination
}
type UpdateBookResponse struct {
	Meta
	Data *Models.Book `json:"data"`
}
type StatusProductResponse struct {
	Meta
}
type GetBookByIdResponse struct {
	Meta
	Data *Models.Book `json:"data"`
}

type Meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}
