package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type NewUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Username *string `json:"username"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
