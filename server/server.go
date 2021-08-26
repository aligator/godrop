package server

import (
	"context"
	"github.com/aligator/godrop"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aligator/godrop/server/graph"
	"github.com/aligator/godrop/server/graph/generated"
	"github.com/aligator/godrop/server/provider"
	"github.com/aligator/godrop/server/service"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8080"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repos, err := provider.NewDefaultRepos()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			NodeService: &service.NodeService{
				Repos: repos,
			},
		},
	}))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		log.Println(e.Error())
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.Handle("/schema.graphqls", &godrop.SchemaHandler{})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
