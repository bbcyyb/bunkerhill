package blog_storage

import (
	"github.com/bbcyyb/bunkerhill/storage"
	"gopkg.in/mgo.v2/bson"
)

type Blog struct {
	ID bson.ObjectId `bson:"_id"`

	Title string `bson:"title"`

	Body string `bson:"body"`

	BodyHTML string `bson:"bodyhtml"`

	Timestamp string `bson:"timestamp"`

	CommentIds []string `bson:"commitids"`

	AuthorId string `bson:"authorid"`
}

var (
	collection = "blog"
)

func NewBlog() *Blog {
	return &Blog{}
}

func GetById(id string) (*Blog, error) {
	blog, err := storage.GetById(collection, id)
	return blog.(*Blog), err
}

func GetAll() ([]Blog, error) {
	var blogs []Blog
	docs, err := storage.GetAll(collection)
	for _, doc := range docs {
		blogs = append(blogs, doc.(Blog))
	}
	return blogs, err
}

func Insert(b *Blog) (string, error) {
	return storage.Insert(collection, b)
}

func (b *Blog) Update(change bson.M) error {
	query := bson.M{"_id": b.ID}
	return storage.Update(collection, query, change)
}

func Remove(selector bson.M) error {
	return storage.Remove(collection, selector)
}

func Get(
	query map[string]interface{},
	sort []string,
	fields map[string]interface{},
	skip int,
	limit int) (results []interface{}, err error) {
	return storage.Get(collection, query, sort, fields, skip, limit)
}
