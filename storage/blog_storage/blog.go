package blog_storage

import (
	"errors"
	"fmt"
	"github.com/bbcyyb/bunkerhill/models"
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

func GetById(id string) (*models.Blog, error) {
	source, err := storage.GetById(collection, id)
	b := source.(*Blog)
	return mtos(b), err
}

func GetAll() ([]*models.Blog, error) {
	var blogs []*models.Blog
	docs, err := storage.GetAll(collection)
	for _, doc := range docs {
		b := doc.(*models.Blog)
		blogs = append(blogs, b)
	}
	return blogs, err
}

func Insert(nb *models.Blog) (string, error) {
	b := stom(nb)
	return storage.Insert(collection, b)
}

func Update(id string, b *models.Blog) error {
	change := bson.M{
		"title":      b.Title,
		"body":       b.Body,
		"bodyhtml":   b.BodyHTML,
		"commentids": b.CommentIds,
	}
	if !bson.IsObjectIdHex(id) {
		return errors.New(fmt.Sprint("id [%s] is not a valid hex representation", id))
	}
	query := bson.M{"_id": bson.ObjectIdHex(id)}
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
	limit int) ([]*models.Blog, error) {
	result, err := storage.Get(collection, query, sort, fields, skip, limit)
	if err != nil {
		return nil, err
	}

	var payload []*models.Blog
	for _, source := range result {
		b := source.(Blog)
		payload = append(payload, mtos(&b))
	}

	return payload, nil
}

func stom(source *models.Blog) *Blog {
	result := &Blog{
		ID:         bson.ObjectIdHex(source.ID),
		Title:      source.Title,
		Body:       source.Body,
		BodyHTML:   source.BodyHTML,
		Timestamp:  source.Timestamp,
		CommentIds: source.CommentIds,
		AuthorId:   source.Author.ID,
	}

	return result
}

func mtos(source *Blog) *models.Blog {
	result := &models.Blog{
		ID:         source.ID.Hex(),
		Title:      source.Title,
		Body:       source.Body,
		BodyHTML:   source.BodyHTML,
		Timestamp:  source.Timestamp,
		CommentIds: source.CommentIds,
	}

	return result
}
