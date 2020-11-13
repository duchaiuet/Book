package response

import "Book/Database/Models"

type CategoryResponse struct {
	Category *Models.Category
	Meta
}
type ListCategoryResponse struct {
	Category []*Models.Category
	Meta
}
