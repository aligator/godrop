package server

import (
	"context"
	"errors"
	"github.com/aligator/godrop"
	"github.com/aligator/godrop/server/file"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

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

func Run(frontend fs.FS) {
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

	if frontend != nil {
		httpFS := http.FileServer(http.FS(frontend))
		router.Handle("/*", http.StripPrefix("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer httpFS.ServeHTTP(w, r)

			if r.URL.Path == "" {
				r.URL.Path = "/"
				return
			}

			f, err := frontend.Open(strings.TrimSuffix(r.URL.Path, "/"))
			if errors.Is(err, fs.ErrNotExist) {
				r.URL.Path = "/"
				return
			} else if err != nil {
				panic(err)
			} else {
				err := f.Close()
				if err != nil {
					panic(err)
				}
			}
		})))
	}
	router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.Handle("/schema.graphql", &godrop.SchemaHandler{})
	router.Handle("/file/*", &file.Handler{
		FileService: &service.FileService{
			Repos: repos,
		},
		TrimSuffix: "/file",
	})

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
