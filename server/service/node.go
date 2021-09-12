package service

import (
	"context"
	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/log"

	"github.com/aligator/godrop/server/provider"
	"github.com/aligator/godrop/server/repository/model"
)

type FileNodeService struct {
	Logger log.GoDropLogger
	Repos  *provider.Repositories
}

func (n FileNodeService) GetFileNodeByPath(ctx context.Context, path string) (model.FileNode, error) {
	return n.Repos.Node.GetByPath(ctx, path)
}

func (n FileNodeService) CreateFileNode(ctx context.Context, newFileNode model.CreateFileNode) (model.FileNode, error) {
	return n.Repos.Node.Create(ctx, newFileNode)
}

func (n FileNodeService) DeleteFileNode(ctx context.Context, id model.ID) (model.ID, error) {
	// 1. Delete node.
	err := n.Repos.Node.DeleteById(ctx, id)
	if err != nil {
		return "", checkpoint.From(err)
	}

	// 2. Delete actual file.
	err = n.Repos.File.Delete(ctx, id)
	if err != nil {
		return "", checkpoint.From(err)
	}

	return id, nil
}
