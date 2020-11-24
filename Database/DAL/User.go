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

type userRepository struct {
	UserCollection   *mongo.Collection
}

func (u userRepository) UpdateStatus(id string, status bool) error {
	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": newId,
	}
	update := bson.M{"$set": bson.M{"status": status}}

	result, err := u.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return  err
	}
	if result.MatchedCount == 0 {
		return errors.New("not update")
	}
	return  err
}

func (u userRepository) Create(user Models.User) (*Models.User, error) {
	user.Id = primitive.NewObjectID()
	result, err := u.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	var userResult *Models.User
	err = u.UserCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&userResult)
	if err != nil {
		return nil, err
	}
	return userResult, err
}

func (u userRepository) Get(id string) (*Models.User, error) {
	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user *Models.User
	err = u.UserCollection.FindOne(context.TODO(), bson.M{"_id": newId}).Decode(&user)
	if err != nil {
		return nil,err
	}
	return user, err
}

func (u userRepository) Gets() ([]*Models.User, error) {
	var users []*Models.User
	cur, err := u.UserCollection.Find(context.TODO(), bson.M{})
	if err != nil{
		return nil, err
	}
	for cur.Next(context.TODO()) {


		var elem *Models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, elem)
	}

	return users, err
}

func (u userRepository) Update(id string, user Models.User) (*Models.User, error) {
	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"_id": newId,
	}
	update := bson.M{"$set": user}

	result, err := u.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	var userResult *Models.User
	err = u.UserCollection.FindOne(context.TODO(), bson.M{"_id": result.UpsertedID}).Decode(&userResult)
	if err != nil {
		return nil, err
	}
	return userResult, err
}

func NewUserRepository(store *Database.MongoDBStore) Models.UserRepository{
	UserCollection := store.Db.Collection(Database.UserCollection)
	return userRepository{
		UserCollection: UserCollection,

	}
}