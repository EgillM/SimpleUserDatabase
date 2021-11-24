package database

import (
	"context"
	"log"
	"time"

	"github.com/EgillM/SimpleUserDatabase/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		client: client,
	}
}

func (db *DB) NewUser(input *model.NewUser) *model.User {
	collection := db.client.Database("userdatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//TODO: Check that username is not present

	//TODO: Hash the password before inserting

	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.User{
		ID:       res.InsertedID.(primitive.ObjectID),
		Username: input.Username,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}

func (db *DB) ListAllUsers() []*model.User {
	collection := db.client.Database("userdatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var users []*model.User
	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}

func (db *DB) FindUserByID(ID string) *model.User {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("userdatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	user := model.User{}
	res.Decode(&user)
	return &user
}

func (db *DB) EraseUser(ID string) *model.User {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("userdatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	res := collection.FindOneAndDelete(ctx, bson.M{"_id": ObjectID})
	user := model.User{}
	res.Decode(&user)
	return &user
}

func (db *DB) Update(ID string, input *model.UpdateUser) *model.User {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("userdatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//TODO: Hash password if it is updated.

	update := bson.D{{"$set",
		bson.D{
			{Key: "username", Value: input.Username},
			{Key: "name", Value: input.Name},
			{Key: "email", Value: input.Email},
			{Key: "password", Value: input.Password},
		}}}
	res := collection.FindOneAndUpdate(ctx, bson.M{"_id": ObjectID}, update)
	user := model.User{}
	res.Decode(&user)
	return &user
}
