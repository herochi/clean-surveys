package model

type SurveyAnswered struct {
	SurveyID   string   `bson:"surveyname" json:"nombre"`
	AnsweredBy string   `bson:"createdBy" json:"userID"`
	Latitude   float32  `bson:"latitude" json:"latitud"`
	Longitude  float32  `bson:"longitude" json:"longitud"`
	Answers    []string `bson:"questions" json:"preguntas"`
}
