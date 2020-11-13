package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
	CategoryIDs []primitive.ObjectID `json:"category_ids" bson:"category_ids"`
	RelatedBookIds []primitive.ObjectID `json:"related_book_ids" bson:"related_book_ids"`
	Images []string `json:"images" bson:"images"`
	Price float32 `json:"price" bson:"price"`
	CommentId []primitive.ObjectID `json:"comment_id" bson:"comment_id"`
	Author string `json:"author" bson:"author"`
	Status bool `json:"status" bson:"status"`
	Quantity int32 `json:"quantity" bson:"quantity"`
}

type BookRepository interface {
	Create(book Book) (*Book, error)
	GetById(id primitive.ObjectID) (*Book, error)
	Filter(code string, author string, status bool, category []primitive.ObjectID, page int, pageSize int ) ([]*Book, int64, error)
	Update(id primitive.ObjectID, book Book) (*Book, error)
	UpdateStatus(status bool, id primitive.ObjectID) error
}