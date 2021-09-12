package provider

import (
	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/provider/filesystem"
	"github.com/aligator/godrop/server/repository"
	"github.com/spf13/afero"
	"os"
)

type Repositories struct {
	Node repository.FileNode
	File repository.File
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

	repos := &Repositories{
		Node: nodeProvider,
		File: fileProvider,
	}

	return repos, nil
}
