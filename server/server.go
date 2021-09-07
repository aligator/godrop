package server

import (
	"context"
	"github.com/aligator/godrop"
	"github.com/aligator/godrop/server/file"
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
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8080"

func Run() {
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
	}).Handler)

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
			FileNodeService: &service.FileNodeService{
				Repos: repos,
			},
		},
	}))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		log.Println(e.Error())
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.Handle("/schema.graphql", &godrop.SchemaHandler{})
	router.Handle("/file/*", &file.Handler{
		FileService: &service.FileService{
			Repos: repos,
		},
		TrimSuffix: "/file",
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
