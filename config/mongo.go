package config

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func MongoConfig() (*mgo.Database, *mgo.Session, error) {

	// uri := os.Getenv("MLAB_URI")
	uri := "mongodb://fidellr:science97@ds239930.mlab.com:39930/kopi-isi-v1"
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	// dbname := os.Getenv("MLAB_DB")
	dbname := "kopi-isi-v1"
	if dbname == "" {
		fmt.Println("no database string provided")
		os.Exit(1)
	}

	session, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo %v\n", err)
		os.Exit(1)
		return nil, session, err
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})

	db := session.DB(dbname)

	return db, session, nil

}
