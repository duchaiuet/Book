package DAL

import (
	"Book/Database"
	"Book/Database/Models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type categoryRepository struct {
	CategoryCollection   *mongo.Collection
}

func (c categoryRepository) Create(category Models.Category) (*Models.Category, error) {
	category.Id = primitive.NewObjectID();
	result, err := c.CategoryCollection.InsertOne(context.TODO(), category)
	if err != nil {
		return nil, err
	}
	var cate *Models.Category
	err = c.CategoryCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&cate)
	return cate, err
}

func (c categoryRepository) Gets() ([]*Models.Category, error) {
	cur, err := c.CategoryCollection.Find(context.TODO(), bson.M{"status": true})
	if err != nil {
		return nil, err
	}
	var results []*Models.Category
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem *Models.Category
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	_ = cur.Close(context.TODO())
	return results, err
}

func (c categoryRepository) Update(category Models.Category) (cate *Models.Category,err error) {
	update :=  bson.M{"$set": bson.M{
		"status": category.Status,
		"name": category.Name,
	}}
	cur, err := c.CategoryCollection.UpdateOne(context.TODO(), bson.M{"_id": category.Id}, update)
	if err != nil {
		return nil, err
	}
	if cur.MatchedCount == 0 {
		return nil, errors.New("not update")
	}
	err =  c.CategoryCollection.FindOne(context.TODO(), bson.M{"_id": category.Id}).Decode(&cate)
	if err != nil {
		return nil, err
	}
	return cate, err
}

func NewCategoryRepository(store *Database.MongoDBStore) Models.CategoryRepository{
	CategoryCollection := store.Db.Collection(Database.CategoryCollection)
	return categoryRepository{
		CategoryCollection:    CategoryCollection,

	}
}