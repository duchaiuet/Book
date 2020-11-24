package response

import "Book/Database/Models"

type CreateUserResponse struct {
	Meta
	Data *Models.User `json:"data"`
}
type ListUserResponse struct {
	Meta
	Data []*Models.User `json:"data"`
}
type UpdateUserResponse struct {
	Meta
	Data *Models.User `json:"data"`
}

type GetUserByIdResponse struct {
	Meta
	Data *Models.User `json:"data"`
}

