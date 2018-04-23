package blog_storage

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/bbcyyb/bunkerhill/models"
	"github.com/bbcyyb/bunkerhill/storage"
	"gopkg.in/mgo.v2/bson"
)

type Blog struct {
	ID bson.ObjectId `bson:"_id,omitempty"`

	Title string `bson:"title"`

	Body string `bson:"body"`

	BodyHTML string `bson:"bodyhtml"`

	CommentIds []bson.ObjectId `bson:"comment_ids"`

	//AuthorId bson.ObjectId `bson:"authorid"`

	CreatedAt time.Time `bson:"created_at"`

	ModifiedAt time.Time `bson:"modified_at"`
}

type Comment struct {
	ID bson.ObjectId `bson:"_id,omitempty"`

	Body string `bson:"body"`

	BodyHTML string `bson:"bodyhtml"`

	Disabled bool `bson:"disabled"`

	AuthorId bson.ObjectId `bson:"author_id"`

	CreatedAt time.Time `bson:"created_at"`

	ModifiedAt time.Time `bson:"modified_at"`
}

var (
	collection = "blog"
)

func GetById(id string) (*models.Blog, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New(fmt.Sprintf("id{%s} is not a valid hex representation", id))
	}

	source, err := storage.GetById(collection, bson.ObjectIdHex(id))
	if source == nil || err != nil {
		return nil, err
	}

	b := convertInterfaceToBlogStruct(source)
	return mtos(b), nil
}

func Insert(nb *models.Blog) (string, error) {
	b := stom(nb)
	b.ID = bson.NewObjectId()
	timeNow := time.Now()
	b.CreatedAt = timeNow
	b.ModifiedAt = timeNow

	/*
		b.AuthorId = b.ID
	*/
	return b.ID.Hex(), storage.Insert(collection, b)
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

func Remove(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New(fmt.Sprint("id [%s] is not a valid hex representation", id))
	}

	selector := bson.M{"_id": bson.ObjectIdHex(id)}
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
		b := convertInterfaceToBlogStruct(&source)
		payload = append(payload, mtos(b))
	}

	return payload, nil
}

func convertInterfaceToBlogStruct(source *interface{}) *Blog {
	var b Blog
	bsonBytes, _ := bson.Marshal(*source)
	bson.Unmarshal(bsonBytes, &b)
	return &b
}

func stom(source *models.Blog) *Blog {
	result := &Blog{
		Title:    source.Title,
		Body:     source.Body,
		BodyHTML: source.BodyHTML,
	}

	if source.ID != "" {
		result.ID = bson.ObjectIdHex(source.ID)
	}

	if source.CreatedAt != "" {
		if i64, err := strconv.ParseInt(source.CreatedAt, 10, 64); err == nil {
			// time.time <- string
			result.CreatedAt = time.Unix(i64, 0)
		}
	}

	if source.ModifiedAt != "" {
		if i64, err := strconv.ParseInt(source.ModifiedAt, 10, 64); err == nil {
			// time.time <- string
			result.ModifiedAt = time.Unix(i64, 0)
		}
	}

	if source.CommentIds == nil {
		result.CommentIds = make([]bson.ObjectId, 0, 1)
	}

	/*
		if source.Author != nil && source.Author.ID != "" {
			result.AuthorId = bson.ObjectIdHex(source.Author.ID)
		}

		if source.CommentIds != nil && 0 < len(source.CommentIds) {
			array := make([]bson.ObjectId, 0, len(source.CommentIds))
			for _, cid := range source.CommentIds {
				array = append(array, bson.ObjectIdHex(cid))
			}

			result.CommentIds = array
		}
	*/

	return result
}

func mtos(source *Blog) *models.Blog {
	result := &models.Blog{
		ID:         source.ID.Hex(),
		Title:      source.Title,
		Body:       source.Body,
		BodyHTML:   source.BodyHTML,
		CreatedAt:  strconv.FormatInt(source.CreatedAt.Unix(), 10),
		ModifiedAt: strconv.FormatInt(source.ModifiedAt.Unix(), 10),
	}

	/*
		if source.CommentIds != nil && 0 < len(source.CommentIds) {
			array := make([]string, 0, len(source.CommentIds))
			for _, cid := range source.CommentIds {
				array = append(array, cid.Hex())
			}

			result.CommentIds = array
		}
	*/

	return result
}
