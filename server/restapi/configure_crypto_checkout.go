// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/rs/cors"

	"github.com/crypto-checkout/handlers"
	"github.com/crypto-checkout/server/restapi/operations"
)

//go:generate swagger generate server --target ../../server --name CryptoCheckout --spec ../../spec/combined_spec.yml --principal interface{}

func configureFlags(api *operations.CryptoCheckoutAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CryptoCheckoutAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	fmt.Println("implementing api in configure api")
	handlers.ImplementAPI(api)

	//api.PreServerShutdown = func() {}

	// api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handler = addCORS(handler)
	return handler
}

func addCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("im in cors")
		c := cors.New(cors.Options{

			AllowedOrigins: []string{"http://localhost:3000", "*"},
			AllowedMethods: []string{"POST", "OPTIONS", "GET", "PUT", "PATCH"},
			AllowedHeaders: []string{
				"Access-Control-Allow-Origin",
				"Access-Control-Allow-Credentials",
				"Accept",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"X-CSRF-Token",
				"Authorization",
			},
			AllowCredentials: true,
			// Enable Debugging for testing, consider disabling in production
			Debug:                true,
			OptionsSuccessStatus: 200,
		})
		handleCORS := c.Handler
		handleCORS(handler).ServeHTTP(res, req)
	})
}
