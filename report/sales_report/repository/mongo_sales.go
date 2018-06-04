package repository

import (
	"time"

	"github.com/sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"

	"github.com/kopi-isi-api-v1/model/report"
	mgo "gopkg.in/mgo.v2"
)

type SalesReportMongo struct {
	db         *mgo.Database
	collection string
}

func NewSalesReportMongo(db *mgo.Database, collection string) *SalesReportMongo {
	return &SalesReportMongo{
		db:         db,
		collection: collection,
	}
}

func (r *SalesReportMongo) Save(srPayload *report.SalesReportEntity) (*report.SalesReportEntity, error) {
	srPayload.CreatedAt = time.Now()
	srPayload.UpdatedAt = time.Now()

	err := r.db.C(r.collection).Insert(srPayload)
	if err != nil {
		return nil, err
	}

	return srPayload, nil
}

func (r *SalesReportMongo) FindByID(id string) (*report.SalesReportEntity, error) {
	srStruct := new(report.SalesReportEntity)
	idBson := bson.ObjectIdHex(id)
	err := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(srStruct)
	if err != nil {
		return nil, err
	}

	return srStruct, nil
}

func (r *SalesReportMongo) FindAll() ([]*report.SalesReportEntity, error) {
	var listOfSr []*report.SalesReportEntity
	err := r.db.C(r.collection).Find(bson.M{}).All(&listOfSr)
	if err != nil {
		return nil, err
	}

	return listOfSr, nil
}

func (r *SalesReportMongo) Update(id string, srPayload *report.SalesReportEntity) (*report.SalesReportEntity, error) {
	srStruct := new(report.SalesReportEntity)
	idBson := bson.ObjectIdHex(id)

	errFind := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(srStruct)
	if errFind != nil {
		logrus.Error(errFind)
	} else if errFind == nil {
		srPayload.CreatedAt = srStruct.CreatedAt
		srPayload.UpdatedAt = time.Now()
		errUpd := r.db.C(r.collection).Update(bson.M{"_id": idBson}, srPayload)
		if errUpd != nil {
			return nil, errUpd
		}
	}

	return srPayload, nil
}

func (r *SalesReportMongo) Delete(id string) (bool, error) {
	srStruct := new(report.SalesReportEntity)
	idBson := bson.ObjectIdHex(id)

	errFind := r.db.C(r.collection).Find(bson.M{"_id": idBson}).One(srStruct)
	if errFind != nil {
		logrus.Error(errFind)
	} else if errFind == nil {
		errRem := r.db.C(r.collection).Remove(bson.M{"_id": idBson})
		if errRem != nil {
			return false, errRem
		}
	}

	return true, nil
}
