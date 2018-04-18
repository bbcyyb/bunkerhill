package blog_imp

import (
	"strings"

	"github.com/bbcyyb/bunkerhill/models"
	"github.com/bbcyyb/bunkerhill/restapi/operations/blog"
	"github.com/bbcyyb/bunkerhill/storage/blog_storage"
	middleware "github.com/go-openapi/runtime/middleware"
)

/*
Restful API 设计规范实战
https://segmentfault.com/a/1190000007313505
*/

func GetBlogs(params blog.GetBlogsParams) middleware.Responder {
	query := make(map[string]interface{})
	select_ := make(map[string]interface{})
	prePage := int(*params.PrePage)
	page := int(*params.Page)
	var sort []string
	var skip, limit int

	if *params.AuthorID != "" {
		query["authorid"] = *params.AuthorID
	}

	if *params.Sortby != "" {
		for _, fragment := range strings.Split(*params.Sortby, ",") {
			var field string
			if fragment[0] == 43 {
				field = fragment[1:]
			} else {
				field = fragment
			}

			sort = append(sort, field)
		}
	}

	if *params.Select != "" {
		for _, field := range strings.Split(*params.Select, ",") {
			select_[field] = 1
		}
	}

	if prePage != 0 && page != 0 {
		skip = (page - 1) * prePage
		limit = prePage
	}

	blogsSource, err := blog_storage.Get(query, sort, select_, skip, limit)

	if err != nil {
		err_payload := &models.GenericError{
			Message: err.Error(),
		}

		return blog.NewGetBlogsInternalServerError().WithPayload(err_payload)
	}

	var payload models.Blogs
	for _, source := range blogsSource {
		payload = append(payload, transfer(source.(blog_storage.Blog)))
	}

	return blog.NewGetBlogsOK().WithPayload(payload)
}

func InsertBlog(params blog.InsertBlogParams) middleware.Responder {
	payload := &models.Blog{}
	return blog.NewInsertBlogCreated().WithPayload(payload)
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
