package blog_imp

import (
	"log"
	"strings"

	"github.com/bbcyyb/bunkerhill/handlers/user_imp"
	"github.com/bbcyyb/bunkerhill/models"
	"github.com/bbcyyb/bunkerhill/restapi/operations/blog"
	"github.com/bbcyyb/bunkerhill/storage/blog_storage"
	middleware "github.com/go-openapi/runtime/middleware"
)

/*
Restful API 设计规范实战
https://segmentfault.com/a/1190000007313505
*/

func Get(params blog.GetBlogsParams) middleware.Responder {
	query := make(map[string]interface{})
	select_ := make(map[string]interface{})
	var sort []string
	var skip, limit int
	var prePage, page int

	if params.AuthorID != nil {
		query["authorid"] = *params.AuthorID
	}

	if params.Sortby != nil {
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

	if params.Select != nil {
		for _, field := range strings.Split(*params.Select, ",") {
			select_[field] = 1
		}
	}

	if params.PrePage != nil {
		prePage = int(*params.PrePage)
	}

	if params.Page != nil {
		page = int(*params.Page)
	}

	if prePage != 0 && page != 0 {
		skip = (page - 1) * prePage
		limit = prePage
	}

	payload, err := blog_storage.Get(query, sort, select_, skip, limit)

	if err != nil {
		errPayload := generateErrorPayload(err)
		return blog.NewGetBlogsInternalServerError().WithPayload(errPayload)
	}

	return blog.NewGetBlogsOK().WithPayload(payload)
}

func GetById(params blog.GetBlogByIDParams) middleware.Responder {
	payload, err := blog_storage.GetById(params.BlogID)
	if err != nil {
		errPayload := generateErrorPayload(err)
		return blog.NewGetBlogByIDInternalServerError().WithPayload(errPayload)
	}

	return blog.NewGetBlogByIDOK().WithPayload(payload)
}

func Insert(params blog.InsertBlogParams) middleware.Responder {
	newId, err := blog_storage.Insert(params.Blog)
	log.Println("=======> 1")
	if err != nil {
		log.Println("=======> 2")
		errPayload := generateErrorPayload(err)
		return blog.NewInsertBlogInternalServerError().WithPayload(errPayload)
	}

	//reload blog entity which just has been created.
	log.Println("=======> 3")
	blog_, err := blog_storage.GetById(newId)
	if err != nil {
		log.Println("=======> 4")
		errPayload := generateErrorPayload(err)
		return blog.NewInsertBlogInternalServerError().WithPayload(errPayload)
	}

	return blog.NewInsertBlogCreated().WithPayload(blog_)
}

func Update(params blog.UpdateBlogParams) middleware.Responder {
	id := params.BlogID
	b := params.Blog
	var author *models.User

	if b.Author != nil {
		a, err := user_imp.UpdateUserCore(b.Author.ID, b.Author)
		if err != nil {
			errPayload := generateErrorPayload(err)
			return blog.NewUpdateBlogInternalServerError().WithPayload(errPayload)
		}
		author = a
	}

	if err := blog_storage.Update(id, b); err != nil {
		errPayload := generateErrorPayload(err)
		return blog.NewUpdateBlogInternalServerError().WithPayload(errPayload)
	}

	newBlog, err := blog_storage.GetById(id)
	if err != nil {
		errPayload := generateErrorPayload(err)
		return blog.NewUpdateBlogInternalServerError().WithPayload(errPayload)
	}

	payload := newBlog
	payload.Author = author
	return blog.NewUpdateBlogOK().WithPayload(payload)
}

func Delete(params blog.DeleteBlogParams) middleware.Responder {
	if err := blog_storage.Remove(params.BlogID); err != nil {
		errPayload := generateErrorPayload(err)
		return blog.NewDeleteBlogInternalServerError().WithPayload(errPayload)
	}

	return blog.NewDeleteBlogOK()
}

func generateErrorPayload(err error) *models.GenericError {
	return &models.GenericError{
		Message: err.Error(),
	}
}
