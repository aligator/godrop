package service

import (
	"context"

	"github.com/aligator/godrop/server/provider"
	"github.com/aligator/godrop/server/repository/model"
)

type NodeService struct {
	Repos *provider.Repositories
}

func (n NodeService) GetNode(ctx context.Context, path string) (model.Node, error) {
	return n.Repos.Note.GetNode(ctx, path)
}

func (n NodeService) CreateNode(ctx context.Context, newNode model.CreateNode) (model.Node, error) {
	return n.Repos.Note.CreateNode(ctx, newNode)
}
