package DAL

import (
	"Book/Database"
	"Book/Database/Models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type bookRepository struct {
	BookCollection   *mongo.Collection
}

func (b bookRepository) Filter(text string, author string, status bool, category []primitive.ObjectID, page int, pageSize int) ([]*Models.Book, int64, error) {
	findOptions := options.Find().SetLimit(int64(pageSize)).SetSkip(int64((page - 1) * pageSize))
	opts := make([]bson.M, 0)
	opts = append(opts, bson.M{"code": primitive.Regex{Pattern: text, Options: "i"}}, bson.M{"name": primitive.Regex{Pattern: text, Options: "i"}})
	q := bson.M{}
	q["$or"] = opts
	books := make([]*Models.Book, 0)
	if page < 0 || pageSize < 0 {
		cur, err := b.BookCollection.Find(context.TODO(), q)
		if err != nil {
			return nil, 0, err
		}
		for cur.Next(context.TODO()) {
			var item *Models.Book
			if err = cur.Decode(&item); err != nil {
				return nil, 0, err
			}
			books = append(books, item)
		}
		_ = cur.Close(context.TODO())
	} else {
		cur, err := b.BookCollection.Find(context.TODO(), q, findOptions)
		for cur.Next(context.TODO()) {
			var item *Models.Book
			if err = cur.Decode(&item); err != nil {
				return nil, 0, err
			}
			books = append(books, item)
		}
		_ = cur.Close(context.TODO())
	}

	totalResult, err := b.BookCollection.CountDocuments(context.TODO(), q)
	return books, totalResult, err
}

func (b bookRepository) Create(book Models.Book) (*Models.Book, error) {
	book.Id = primitive.NewObjectID()
	result, err := b.BookCollection.InsertOne(context.TODO(), book)
	if err != nil {
		return nil, err
	}
	var temp *Models.Book
	err = b.BookCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&temp)
	if err != nil {
		return nil, err
	}
	return temp, err
	
}

func (b bookRepository) GetById(id primitive.ObjectID) (*Models.Book, error) {
	var book *Models.Book
	err := b.BookCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return book, err
}


func (b bookRepository) Update(id primitive.ObjectID, book Models.Book) (*Models.Book, error) {
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{"$set": book}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	// 8) Find one result and update it
	var result *Models.Book
	err :=  b.BookCollection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (b bookRepository) UpdateStatus(status bool, id primitive.ObjectID) error {
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{"$set": bson.M{
		"status": status,
	}}
	result, err := b.BookCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0{
		return errors.New("not updated")
	}
	return err
}

func NewBookRepository(store *Database.MongoDBStore) Models.BookRepository{
	BookCollection := store.Db.Collection(Database.BookCollection)
	return bookRepository{
		BookCollection:    BookCollection,

	}
}