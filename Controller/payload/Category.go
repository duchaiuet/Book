package payload

type Category struct {
	Status bool `json:"status" bson:"status"`
	Name string `json:"name" bson:"name"`
}