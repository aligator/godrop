package main

import (
	"github.com/aligator/godrop/server"
	"github.com/aligator/godrop/server/log"
	"github.com/aligator/godrop/server/provider"
	"os"
)

func main() {
	logger := log.DefaultLogger()

	filesLocation := os.Getenv("GODROP_FILES")
	if filesLocation == "" {
		filesLocation = "./files"
	}

	repos, err := provider.NewDefaultRepos(filesLocation)
	if err != nil {
		logger.Fatal(err)
	}

	s := server.Server{
		Host:           os.Getenv("GODROP_HOST"),
		AllowedOrigins: []string{"*"},
		Repos:          repos,
		Logger:         logger,
	}

	if s.Host == "" {
		// Set localhost as default else windows needs to open the firewall
		// for each new compiled binary, which is very annoying in development.
		s.Host = "localhost:8080"
	}

	_ = s.Init()
	s.Run()
}
