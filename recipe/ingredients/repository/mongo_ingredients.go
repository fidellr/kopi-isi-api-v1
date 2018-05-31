package repository

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/kopi-isi-api-v1/model/recipe"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IngredientsMongo struct {
	db         *mgo.Database
	collection string
}

func NewIngredientsMongo(db *mgo.Database, collection string) *IngredientsMongo {
	return &IngredientsMongo{
		db:         db,
		collection: collection,
	}
}

func (r *IngredientsMongo) Save(ingPayload *recipe.IngredientMaster) (*recipe.IngredientMaster, map[string]string, error) {
	ingStruct := new(recipe.IngredientMaster)
	ingPayload.CreatedAt = time.Now()
	ingPayload.UpdatedAt = time.Now()

	errFind := r.db.C(r.collection).Find(bson.M{"name": ingPayload.Name}).One(ingStruct)
	if errFind != nil {
		fmt.Println("errFind ing save err:", errFind)
	} else if errFind == nil {
		return nil, map[string]string{
			"name": "ingridient is exist",
		}, nil
	}

	errIns := r.db.C(r.collection).Insert(ingPayload)
	if errIns != nil {
		return nil, nil, errIns
	}

	return ingPayload, nil, nil
}

func (r *IngredientsMongo) FindByID(id string) (*recipe.IngredientMaster, error) {
	ingStruct := new(recipe.IngredientMaster)
	idBson := bson.ObjectIdHex(id)
	err := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(ingStruct)
	if err != nil {
		return nil, err
	}

	return ingStruct, nil
}

func (r *IngredientsMongo) FindAll() ([]*recipe.IngredientMaster, error) {
	var listOfIng []*recipe.IngredientMaster
	err := r.db.C(r.collection).Find(bson.M{}).All(&listOfIng)
	if err != nil {
		return nil, err
	}

	return listOfIng, nil
}

func (r *IngredientsMongo) Update(id string, ingPayload *recipe.IngredientUpdate) (*recipe.IngredientUpdate, error) {
	ingStruct := new(recipe.IngredientMaster)
	idBson := bson.ObjectIdHex(id)

	errFind := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(ingStruct)
	if errFind != nil {

	} else if errFind == nil {
		ingPayload.CreatedAt = ingStruct.CreatedAt
		ingPayload.UpdatedAt = time.Now()
		errUpd := r.db.C(r.collection).Update(bson.M{"_id": idBson}, ingPayload)
		if errUpd != nil {
			return nil, errUpd
		}
	}

	return ingPayload, nil
}

func (r *IngredientsMongo) Delete(id string) (bool, error) {
	ing := new(recipe.IngredientMaster)
	idBson := bson.ObjectIdHex(id)

	errFind := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(ing)
	if errFind != nil {
		logrus.Error(errFind)
	} else if errFind == nil {
		err := r.db.C(r.collection).Remove(bson.M{"_id": idBson})
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
