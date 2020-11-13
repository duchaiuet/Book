package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Status bool `json:"status" bson:"status"`
	Name string `json:"name" bson:"name"`
}

type CategoryRepository interface {
	Create(category Category) (*Category, error)
	Gets() ([]*Category, error)
	Update(category Category) (*Category, error)
}
