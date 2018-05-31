package repository

import (
	"fmt"
	"time"

	"github.com/kopi-isi-api-v1/model/recipe"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CoffeeMongo struct {
	db         *mgo.Database
	collection string
}

func NewCoffeeMongo(db *mgo.Database, collection string) *CoffeeMongo {
	return &CoffeeMongo{
		db:         db,
		collection: collection,
	}
}

func (r *CoffeeMongo) Save(cofPayload *recipe.CoffeeEntity) (*recipe.CoffeeEntity, map[string]string, error) {
	cofStruct := new(recipe.CoffeeEntity)
	cofPayload.CreatedAt = time.Now()
	cofPayload.UpdatedAt = time.Now()

	errFind := r.db.C(r.collection).Find(bson.M{"coffee_name": cofPayload.CoffeeName}).One(cofStruct)
	if errFind != nil {
		// fmt.Println("errFind cof save err:", errFind)
	} else if errFind == nil {
		return nil, map[string]string{
			"coffee_name": "coffee name is exist",
		}, nil
	}

	errIns := r.db.C(r.collection).Insert(cofPayload)
	if errIns != nil {
		fmt.Println("errIns coffe err:", errIns)
		return nil, nil, errIns
	}

	return cofPayload, nil, nil
}

func (r *CoffeeMongo) Update(id string, cofPayload *recipe.CoffeeEntity) (*recipe.CoffeeEntity, error) {
	cofStruct := new(recipe.CoffeeEntity)
	idBson := bson.ObjectIdHex(id)

	errFind := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(cofStruct)
	if errFind != nil {
		return nil, errFind
	}

	cofPayload.CreatedAt = cofStruct.CreatedAt
	cofPayload.UpdatedAt = time.Now()
	errUpd := r.db.C(r.collection).Update(bson.M{"_id": idBson}, cofPayload)
	if errUpd != nil {
		return nil, errUpd
	}

	return cofPayload, nil
}

func (r *CoffeeMongo) FindByID(id string) (*recipe.CoffeeEntity, error) {
	cofStruct := new(recipe.CoffeeEntity)
	idBson := bson.ObjectIdHex(id)
	err := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(cofStruct)
	if err != nil {
		return nil, err
	}

	return cofStruct, nil
}

func (r *CoffeeMongo) FindAll() ([]*recipe.CoffeeEntity, error) {
	var listOfCof []*recipe.CoffeeEntity
	err := r.db.C(r.collection).Find(bson.M{}).All(&listOfCof)
	if err != nil {
		return nil, err
	}

	return listOfCof, nil
}

func (r *CoffeeMongo) Delete(id string) (bool, error) {
	coffee := new(recipe.CoffeeEntity)
	idBson := bson.ObjectIdHex(id)

	errFind := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(coffee)
	if errFind != nil {
		return false, errFind
	}

	err := r.db.C(r.collection).Remove(bson.M{"_id": idBson})
	if err != nil {
		return false, err
	}

	return true, nil
}
