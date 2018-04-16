package storage

import (
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

func GetById(id string) (*Blog, error) {
}

func Get() ([]Blog, error) {
}

func Insert(b *Blog) (string, error) {
}

func Update(query bson.M, change bson.M) err {
}

func Remove(selector bson.M) err {

}

func Search(
	query bson.M, sort string,
	fields bson.M,
	skip int,
	limit int) (results []interface{}, err error) {
}
