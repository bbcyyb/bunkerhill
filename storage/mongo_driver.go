package storage

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDoc struct {
	ID bson.ObjectId
}

var (
	mgoSession *mgo.Session
	database   = "bunkerhill"
	URL        = "mongodb://10.62.59.210:27017"
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err)
		}
	}
	return mgoSession.Clone()
}

func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(database).C(collection)
	return s(c)
}

func GetById(collection string, id string) (interface{}, error) {
	objId := bson.ObjectId(id)
	var result interface{}
	operation := func(c *mgo.Collection) error {
		return c.Find(objId).One(&result)
	}

	if err := withCollection(collection, operation); err != nil {
		return nil, err
	}

	return &result, nil
}

func Get(collection string) ([]interface{}, error) {
	var result []interface{}
	operation := func(c *mgo.Collection) error {
		return c.Find(nil).All(&result)
	}

	if err := withCollection(collection, operation); err != nil {
		return nil, err
	}

	return result, nil
}

func Insert(collection string, doc *MongoDoc) (string, error) {
	doc.ID = bson.NewObjectId()
	operation := func(c *mgo.Collection) error {
		return c.Insert(*doc)
	}

	if err := withCollection(collection, operation); err != nil {
		return "", err
	}

	return doc.ID.Hex(), nil
}

func Update(collection string, query bson.M, change bson.M) error {
	operation := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}

	return withCollection(collection, operation)
}

func Remove(collection string, selector bson.M) error {
	if selector == nil || 0 == len(selector) {
		panic("No selector exists ....")
	}

	operation := func(c *mgo.Collection) error {
		return c.Remove(selector)
	}

	return withCollection(collection, operation)
}

func Search(
	collection string,
	query bson.M, sort string,
	fields bson.M,
	skip int,
	limit int) (results []interface{}, err error) {
	operation := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}

	if err = withCollection(collection, operation); err != nil {
		return nil, err
	}

	return results, nil
}
