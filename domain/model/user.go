package model

type User struct {
	Username     string `bson:"username" json:"username"`
	Password     string `bson:"password" json:"password"`
	Email        string `bson:"email" json:"email"`
	Investigator bool
}
