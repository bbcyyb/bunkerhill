package blog_imp

import (
	"github.com/bbcyyb/bunkerhill/models"
	"github.com/bbcyyb/bunkerhill/restapi/operations/blog"
	"github.com/bbcyyb/bunkerhill/storage/blog_storage"
	middleware "github.com/go-openapi/runtime/middleware"
)

/*
Restful API 设计规范实战
https://segmentfault.com/a/1190000007313505
*/

func GetBlog(params blog.GetBlogParams) middleware.Responder {
	blogsSource, err := blog_storage.GetBlogAll()
	if err != nil {
		err_payload := &models.GenericError{
			Message: err.Error(),
		}

		blog.NewGetBlogInternalServerError().WithPayload(err_payload)
	}

	blogsTarget := make(models.Blogs, len(blogsSource))
	for _, source := range blogsSource {
		blogsTarget = append(blogsTarget, transfer(source))
	}
	payload := &models.GetBlogOKBody{
		Data:   blogsTarget,
		Paging: nil,
	}

	return blog.NewGetBlogOK().WithPayload(payload)
}

func transfer(source blog_storage.Blog) *models.Blog {
	result := &models.Blog{
		ID:         source.ID.Hex(),
		Title:      source.Title,
		Body:       source.Body,
		BodyHTML:   source.BodyHTML,
		Timestamp:  source.Timestamp,
		CommentIds: source.CommentIds,
		Author:     nil,
	}

	return result
}
