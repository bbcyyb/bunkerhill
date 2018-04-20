package storage

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDoc struct {
	ID bson.ObjectId
}

var (
	mgoSession *mgo.Session
	database   = "bunkerhill"
	URL        = "mongodb://127.0.0.1:27017"
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			log.Println(err.Error())
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

func GetById(collection string, id bson.ObjectId) (interface{}, error) {
	var result interface{}
	operation := func(c *mgo.Collection) error {
		return c.Find(id).One(&result)
	}

	if err := withCollection(collection, operation); err != nil {
		return nil, err
	}

	return &result, nil
}

func GetAll(collection string) ([]interface{}, error) {
	var result []interface{}
	operation := func(c *mgo.Collection) error {
		return c.Find(nil).All(&result)
	}

	if err := withCollection(collection, operation); err != nil {
		return nil, err
	}

	return result, nil
}

func Get(
	collection string,
	query bson.M,
	sort []string,
	fields bson.M,
	skip int,
	limit int) (results []interface{}, err error) {
	operation := func(c *mgo.Collection) error {
		var q *mgo.Query
		if len(query) > 0 {
			q = c.Find(query)
		} else {
			q = c.Find(nil)
		}

		if len(sort) > 0 {
			q = q.Sort(sort...)
		}

		if len(fields) > 0 {
			q = q.Select(fields)
		}

		if skip > 0 {
			q = q.Skip(skip)
		}

		if limit > 0 {
			q = q.Limit(limit)
		}

		return q.All(&results)
	}

	if err = withCollection(collection, operation); err != nil {
		return nil, err
	}

	return results, nil
}

func Insert(collection string, d interface{}) (string, error) {
	doc := d.(*MongoDoc)
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
