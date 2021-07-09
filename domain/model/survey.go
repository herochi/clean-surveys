package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Survey struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Surveyname string             `bson:"surveyname" json:"surveyname"`
	Questions  []string           `bson:"questions" json:"questions"`
	CreatedBy  string             `bson:"createdBy" json:"createdBy"`
	CreatedAt  *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt  *time.Time         `bson:"updatedAt" json:"updatedAt"`
}
