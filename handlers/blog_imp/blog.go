package blog_imp

import (
	"github.com/bbcyyb/bunkerhill/models"
	"github.com/bbcyyb/bunkerhill/restapi/operations"
	"github.com/bbcyyb/bunkerhill/restapi/operations/blog"
	"github.com/bbcyyb/bunkerhill/storage/blog_storage"
	middleware "github.com/go-openapi/runtime/middleware"
)

func GetBlog(params blog.GetBlogParams) middleware.Responder {
	blogsSource, err := blog_storage.GetBlogAll()
	if err != nil {
		err_payload := &models.GenericError{
			Message: err.Error(),
		}

		apps.NewGetAPIVersionInternalServerError().WithPayload(err_payload)
	}

	blogsTarget := make(models.Blogs, len(blogsSource))
	for _, source := range blogsSource {
		append(blogsTarget, *transfer(source))
	}

	return blog.NewGetBlogOK().WithPayload(blogsTarget)
}

func transfer(source blog_storage.Blog) *models.Blog {
	result := &models.Blog{
		ID:         source.ID,
		Title:      source.Title,
		Body:       source.Body,
		BodyHTML:   source.BodyHTML,
		Timestamp:  source.Timestamp,
		CommentIds: source.CommentIds,
		Author:     nil,
	}

	return result
}
