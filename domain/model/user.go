package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserID       primitive.ObjectID `bson:"userid" json:"userid"`
	Username     string             `bson:"username" json:"username"`
	Password     string             `bson:"password" json:"password"`
	Email        string             `bson:"email" json:"email"`
	Investigator bool
}
