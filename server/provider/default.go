package provider

import (
	"github.com/aligator/godrop/server/provider/filesystem"
	"github.com/aligator/godrop/server/repository"
)

type Repositories struct {
	Node repository.FileNode
	File repository.File
}

func NewDefaultRepos(filesLocation string) (*Repositories, error) {
	fsProvider := new(filesystem.Provider)
	err := fsProvider.Init(filesLocation)
	if err != nil {
		return nil, err
	}

	repos := &Repositories{
		Node: fsProvider,
		File: fsProvider,
	}

	return repos, nil
}
