package storage

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mgoSession *mgo.Session
	dataBase   = "mydb"
	URL        = "mongodb://10.62.59.210:27017"
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err)
		}

		return mgoSession.Clone()
	}
}

func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(database).C(collection)
	return s(c)
}

func GetById(id string) (*Blog, error) {
	objId := bson.ObjectId(id)
	var blog Blog
	query := func(c *mgo.Collection) error {
		return c.Find(objId).One(&blog)
	}

	if err = withCollection("blog", query); err != nil {
		return nil, err
	}

	return &blog, nil
}

func Get() ([]Blog, error) {
	var blogs []Blog
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&blogs)
	}

	err := withCollection("blog", query)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func Insert(b *Blog) (string, error) {
	b.ID = bson.NewObjectId()
	query := func(c *mgo.Collection) error {
		return c.Insert(*b)
	}

	if err := withCollection("blog", query); err != nil {
		return nil, err
	}

	return p.id.Hex(), nil
}

func Update(query bson.M, change bson.M) err {
	exop := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}

	if err := withCollection("blog", exop); err != nil {
		return err
	}

	return err
}

func Search(
	collectionName string,
	query bson.M, sort string,
	fields bson.M,
	skip int,
	limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}

	if err = withCollection(collectionName, exop); err != nil {
		return nil, err
	}

	return results, nil
}
