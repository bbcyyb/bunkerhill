// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	graceful "github.com/tylerb/graceful"

	"github.com/bbcyyb/bunkerhill/handlers/apiversion_imp"
	"github.com/bbcyyb/bunkerhill/handlers/blog_imp"
	"github.com/bbcyyb/bunkerhill/restapi/operations"
	"github.com/bbcyyb/bunkerhill/restapi/operations/apiversion"
	"github.com/bbcyyb/bunkerhill/restapi/operations/blog"
)

//go:generate swagger generate server --target .. --name bunkerhill --spec ../swagger/swagger.yaml --exclude-main

func configureFlags(api *operations.BunkerhillAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BunkerhillAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	/**************************************
	* Apiversion
	**************************************/
	api.ApiversionGetAPIVersionHandler = apiversion.GetAPIVersionHandlerFunc(apiversion_imp.GetAPIVersion)
	/**************************************
	* Blog
	**************************************/
	api.BlogGetBlogsHandler = blog.GetBlogsHandlerFunc(blog_imp.Get)
	api.BlogGetBlogByIDHandler = blog.GetBlogByIDHandlerFunc(blog_imp.GetById)
	api.BlogInsertBlogHandler = blog.InsertBlogHandlerFunc(blog_imp.Insert)
	api.BlogUpdateBlogHandler = blog.UpdateBlogHandlerFunc(blog_imp.Update)
	api.BlogDeleteBlogHandler = blog.DeleteBlogHandlerFunc(blog_imp.Delete)
	/**************************************
	* User
	**************************************/

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
