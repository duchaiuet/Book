package Database

var (
	DbName         string
	BookCollection string
	CategoryCollection string
	UserCollection string
	HttpSwagger    string
	HostPort    string
)

func loadParameters() {
	DbName = "BookShop"
	BookCollection = "Book"
	CategoryCollection = "Category"
	UserCollection = "User"
	HttpSwagger = "http://localhost:1234/api/v1/book/swagger/doc.json"
	HostPort =  ":1234"
}
