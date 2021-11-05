package provider

import (
	"os"

	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/provider/filesystem"
	"github.com/aligator/godrop/server/provider/keycloak"
	"github.com/aligator/godrop/server/repository"
	"github.com/spf13/afero"
)

type Repositories struct {
	Node repository.FileNode
	File repository.File
	Auth repository.Auth
}

func NewDefaultRepos(filesLocation string) (*Repositories, error) {
	if err := os.Mkdir(filesLocation, 0775); !os.IsExist(err) && err != nil {
		return nil, checkpoint.From(err)
	}

	osFS := afero.NewBasePathFs(afero.NewOsFs(), filesLocation)

	nodeProvider, fileProvider, err := filesystem.New(osFS)
	if err != nil {
		return nil, err
	}

	keycloak.New(keycloak.Config{
		Host:     "http://localhost:7080",
		Realm:    "godrop",
		ClientID: "godrop",
	})

	repos := &Repositories{
		Node: nodeProvider,
		File: fileProvider,
	}

	return repos, nil
}
