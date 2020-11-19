package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Status bool `json:"status" bson:"status"`
}
type RoleRepository interface {
	Create(role Role) (*Role, error)
	Get(id string) (*Role, error)
	Gets() ([]*Role, error)
	Update(role Role) (*Role, error)
	
}