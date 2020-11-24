package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	City     string `json:"city"bson:"city"`
	District string `json:"district" bson:"district"`
	Ward     string `json:"ward" bson:"ward"`
	Text     string `json:"text" bson:"text"`
}

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Code        string             `json:"code" bson:"code"`
	Role        primitive.ObjectID `json:"role" bson:"role"`
	UserName    string             `json:"user_name" bson:"user_name"`
	Password    string             `json:"password" bson:"password"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Address     Address            `json:"address" bson:"address"`
	Status      bool               `json:"status" bson:"status"`
}

type UserRepository interface {
	Create(user User) (*User, error)
	Get(id string) (*User, error)
	Gets() ([]*User, error)
	Update(id string, user User) (*User, error)
	UpdateStatus(id string, status bool) error
}
