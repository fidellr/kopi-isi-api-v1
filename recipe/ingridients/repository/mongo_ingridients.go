package repository

import (
	"time"

	"github.com/kopi-isi-api-v1/model/recipe"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IngridientsMongo struct {
	db         *mgo.Database
	collection string
}

func NewIngridientsMongo(db *mgo.Database, collection string) *IngridientsMongo {
	return &IngridientsMongo{
		db:         db,
		collection: collection,
	}
}

func (r *IngridientsMongo) Save(ingPayload *recipe.IngridientMaster) (*recipe.IngridientMaster, map[string]string, error) {
	ingStruct := new(recipe.IngridientMaster)
	ingPayload.CreatedAt = time.Now()
	ingPayload.UpdatedAt = time.Now()

	errFind := r.db.C(r.collection).Find(bson.M{"name": ingPayload.Name}).One(ingStruct)
	if errFind != nil {
		// return nil, nil, errFind
	} else if errFind == nil {
		return nil, map[string]string{
			"name": "ingridient is exist",
		}, nil
	}

	errIns := r.db.C(r.collection).Insert(ingPayload)
	if errIns != nil {
		panic(errIns)
		return nil, nil, errIns
	}

	return ingPayload, nil, nil
}

func (r *IngridientsMongo) FindByID(id string) (*recipe.IngridientMaster, error) {
	ingStruct := new(recipe.IngridientMaster)
	idBson := bson.ObjectIdHex(id)
	err := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(ingStruct)
	if err != nil {
		return nil, err
	}

	return ingStruct, nil
}

func (r *IngridientsMongo) FindAll() ([]*recipe.IngridientMaster, error) {
	var listOfIng []*recipe.IngridientMaster
	err := r.db.C(r.collection).Find(bson.M{}).All(&listOfIng)
	if err != nil {
		return nil, err
	}

	return listOfIng, nil
}

func (r *IngridientsMongo) Update(id string, ingPayload *recipe.IngridientUpdate) (*recipe.IngridientUpdate, error) {
	ingPayload.UpdatedAt = time.Now()
	idBson := bson.ObjectIdHex(id)
	err := r.db.C(r.collection).Update(bson.M{"_id": idBson}, ingPayload)
	if err != nil {
		return nil, err
	}

	return ingPayload, nil
}

func (r *IngridientsMongo) Delete(id string) (bool, error) {
	idBson := bson.ObjectIdHex(id)
	err := r.db.C(r.collection).Remove(bson.M{"_id": idBson})
	if err != nil {
		return false, err
	}

	return true, nil
}
