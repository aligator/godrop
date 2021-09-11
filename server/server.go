package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aligator/godrop"
	"github.com/aligator/godrop/server/file"
	"github.com/aligator/godrop/server/graph"
	"github.com/aligator/godrop/server/graph/generated"
	"github.com/aligator/godrop/server/log"
	"github.com/aligator/godrop/server/provider"
	"github.com/aligator/godrop/server/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
)

type Server struct {
	Host           string
	AllowedOrigins []string
	Repos          *provider.Repositories
	Logger         log.GoDropLogger

	router *chi.Mux
}

func (s *Server) Init() *chi.Mux {
	s.router = chi.NewRouter()

	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:   s.AllowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Logger: s.Logger,
			FileNodeService: &service.FileNodeService{
				Logger: s.Logger,
				Repos:  s.Repos,
			},
		},
	}))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		s.Logger.Error(e.Error())
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})

	s.router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	s.router.Handle("/query", srv)
	s.router.Handle("/schema.graphql", &godrop.SchemaHandler{})
	s.router.Handle("/file/*", &file.Handler{
		Logger: s.Logger,
		FileService: &service.FileService{
			Logger: s.Logger,
			Repos:  s.Repos,
		},
		TrimSuffix: "/file",
	})

	return s.router
}

func (s *Server) Run() {
	if s.router == nil {
		panic("call Init before Run")
	}

	s.Logger.Printf("connect to http://%s/playground for a GraphQL playground", s.Host)
	s.Logger.Fatal(http.ListenAndServe(s.Host, s.router))
}
