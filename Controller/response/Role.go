package response

import "Book/Database/Models"

type RoleResponse struct {
	Data *Models.Role
	Meta
}

type ListRoleResponse struct {
	Data []*Models.Role
	Meta
}
