package viewmodel

type UserVM struct {
	UserID       string `bson:"userid" json:"userid"`
	Username     string `bson:"username" json:"username"`
	Password     string `bson:"password" json:"password"`
	Email        string `bson:"email" json:"email"`
	Investigator bool
}
