package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"ms-workspace/gateway/internal/delivery/graphql/graph"
	"ms-workspace/gateway/internal/delivery/graphql/graph/resolver"
	"net/http"
)

func graphqlHandler() fiber.Handler {
	resolver := &resolver.Resolver{}

	// TODO init dependencies
	// ---------------------------------------

	// ---------------------------------------
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	return func(c *fiber.Ctx) error {
		// cover authentication session of fiber context
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			srv.ServeHTTP(writer, request)
		})(c.Context())
		return nil
	}
}

func playgroundHandler() fiber.Handler {
	srv := playground.Handler("GraphQL", "/query")
	return func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			srv.ServeHTTP(writer, request)
		})(c.Context())
		return nil
	}
}
