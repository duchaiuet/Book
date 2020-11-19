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

type roleRepository struct {
	RoleCollection *mongo.Collection
}

func (r roleRepository) Create(role Models.Role) (*Models.Role, error) {
	role.Id = primitive.NewObjectID()
	result, err := r.RoleCollection.InsertOne(context.TODO(), role)
	if err != nil {
		return nil, err
	}
	var roleResult *Models.Role
	err = r.RoleCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&roleResult)
	if err != nil {
		return nil, err
	}
	return roleResult, err
}

func (r roleRepository) Get(id string) (*Models.Role, error) {
	var roleResult *Models.Role
	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.RoleCollection.FindOne(context.TODO(), bson.M{"_id": newId}).Decode(&roleResult)
	if err != nil {
		return nil, err
	}
	return roleResult, err
}

func (r roleRepository) Gets() ([]*Models.Role, error) {
	cur, err := r.RoleCollection.Find(context.TODO(), bson.M{"status": true})
	if err != nil {
		return nil, err
	}
	var results []*Models.Role
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem *Models.Role
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

func (r roleRepository) Update(role Models.Role) (*Models.Role, error) {
	update :=  bson.M{"$set": bson.M{
		"status": role.Status,
		"name": role.Name,
	}}
	cur, err := r.RoleCollection.UpdateOne(context.TODO(), bson.M{"_id": role.Id}, update)
	if err != nil {
		return nil, err
	}
	if cur.MatchedCount == 0 {
		return nil, errors.New("not update")
	}
	var roleResult *Models.Role
	err =  r.RoleCollection.FindOne(context.TODO(), bson.M{"_id": role.Id}).Decode(&roleResult)
	if err != nil {
		return nil, err
	}
	return roleResult, err
}

func NewRoleRepository(store *Database.MongoDBStore) Models.RoleRepository {
	roleCollection := store.Db.Collection(Database.RoleCollection)
	return roleRepository{
		RoleCollection: roleCollection,
	}
}
