package provider

import (
	"github.com/aligator/godrop/server/provider/filesystem"
	"github.com/aligator/godrop/server/repository"
)

type Repositories struct {
	Note repository.FileNode
}

func NewDefaultRepos() (*Repositories, error) {
	filesystem := new(filesystem.Provider)

	repos := &Repositories{
		Note: filesystem,
	}
	err := filesystem.Init()
	if err != nil {
		return nil, err
	}

	return repos, nil
}
