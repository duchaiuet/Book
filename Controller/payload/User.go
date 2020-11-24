package payload

type User struct {
	Name        string  `json:"name" bson:"name"`
	UserName    string  `json:"user_name" bson:"user_name"`
	Password    string  `json:"password" bson:"password"`
	PhoneNumber string  `json:"phone_number" bson:"phone_number"`
	Address     Address `json:"address" bson:"address"`
	Role        string  `json:"role"`
}

type Address struct {
	City     string `json:"city"bson:"city"`
	District string `json:"district" bson:"district"`
	Ward     string `json:"ward" bson:"ward"`
	Text     string `json:"text" bson:"text"`
}
