package recipe

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type IngredientMaster struct {
	ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt      time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" bson:"updated_at"`
	Name           string        `json:"name" bson:"name" validate:"required"`
	TotalGram      float64       `json:"total_gram" bson:"total_gram" validate:"required"`
	TotalGramPrice float64       `json:"total_gram_price" bson:"total_gram_price"`
}

type IngredientUpdate struct {
	ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt      time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" bson:"updated_at"`
	Name           string        `json:"name" bson:"name"`
	TotalGram      float64       `json:"total_gram" bson:"total_gram"`
	TotalGramPrice float64       `json:"total_gram_price" bson:"total_gram_price"`
}
