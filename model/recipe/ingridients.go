package recipe

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type IngridientMaster struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	Name      string        `json:"name" bson:"name" validate:"required"`
	Gram      int64         `json:"gram" bson:"gram" validate:"required"`
	Price     int64         `json:"price" bson:"price" validate:"required"`
}

type IngridientUpdate struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	Name      string        `json:"name" bson:"name"`
	Gram      int64         `json:"gram" bson:"gram"`
	Price     int64         `json:"price" bson:"price"`
}
