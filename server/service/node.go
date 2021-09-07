package service

import (
	"context"

	"github.com/aligator/godrop/server/provider"
	"github.com/aligator/godrop/server/repository/model"
)

type FileNodeService struct {
	Repos *provider.Repositories
}

func (n FileNodeService) GetFileNodeByPath(ctx context.Context, path string) (model.FileNode, error) {
	return n.Repos.Node.GetFileNodeByPath(ctx, path)
}

func (n FileNodeService) CreateFileNode(ctx context.Context, newFileNode model.CreateFileNode) (model.FileNode, error) {
	return n.Repos.Node.CreateFileNode(ctx, newFileNode)
}
