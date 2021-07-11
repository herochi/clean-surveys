package model

//"go.mongodb.org/mongo-driver/bson/primitive"

type Survey struct {
	Surveyname string   `bson:"surveyname" json:"nombre"`
	Questions  []string `bson:"questions" json:"preguntas"`
	CreatedBy  string   `bson:"createdBy" json:"userId"`
}
