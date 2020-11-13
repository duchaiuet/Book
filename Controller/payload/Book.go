package payload

type CreateBookPayload struct {
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
	CategoryIDs []string `json:"category_ids" bson:"category_ids"`
	Images []string `json:"images" bson:"images"`
	Price float32 `json:"price" bson:"price"`
	Author string `json:"author" bson:"author"`
	Status bool `json:"status" bson:"status"`
	Quantity int32 `json:"quantity" bson:"quantity"`
}

type UpdateBookPayload struct {
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
	CategoryIDs []string `json:"category_ids" bson:"category_ids"`
	Images []string `json:"images" bson:"images"`
	Price float32 `json:"price" bson:"price"`
	Author string `json:"author" bson:"author"`
	Status bool `json:"status" bson:"status"`
	Quantity int32 `json:"quantity" bson:"quantity"`
}