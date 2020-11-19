package payload

type Role struct {
	Name string `json:"name" bson:"name"`
	Status bool `json:"status" bson:"status"`
}