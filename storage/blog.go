package storage

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Blog struct {
	ID bson.ObjectId `bson:"_id"`

	Title string `bson:"title"`

	Body string `bson:"body"`

	BodyHTML string `bson:"bodyhtml"`

	Timestamp string `bson:"timestamp"`

	CommentIds []int32 `bson:"commitids"`

	AuthorId []int32 `bson:"authorid"`

	CreateAt datetime
}

func (b *Blog) NewBlog() *Blog {
	return &Blog{}
}

func (b *Blog) getAll() ([]Blog, error) {
	session, err := mgo.Dial("mongodb://10.62.59.210:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("bunkerhill").C("blog")
	iter := c.Find().Iter()

	blog := Blog{}
	for iter.Next(&blog) {
	}
	return nil, error()
}
