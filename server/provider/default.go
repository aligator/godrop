package provider

import (
	"github.com/aligator/godrop/server/provider/filesystem"
	"github.com/aligator/godrop/server/repository"
)

type Repositories struct {
	Node repository.FileNode
	File repository.File
}

func NewDefaultRepos() (*Repositories, error) {
	fsProvider := new(filesystem.Provider)

	repos := &Repositories{
		Node: fsProvider,
		File: fsProvider,
	}
	err := fsProvider.Init()
	if err != nil {
		return nil, err
	}

	return repos, nil
}
