package Database

var (
	DbName         string
	BookCollection string
	CategoryCollection string
	UserCollection string
	RoleCollection string


	HttpSwagger    string
	HostPort    string
)

func loadParameters() {
	DbName = "BookShops"
	BookCollection = "Books"
	CategoryCollection = "Categories"
	RoleCollection = "Roles"
	UserCollection = "Users"
	HttpSwagger = "http://localhost:1234/api/v1/book/swagger/doc.json"
	HostPort =  ":1234"
}
