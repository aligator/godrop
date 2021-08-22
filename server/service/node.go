package service

import (
	"context"

	"github.com/aligator/godrop/server/repository/model"
)

type NodeService struct {
}

func (n NodeService) GetNode(ctx context.Context, path string) (model.Node, error) {

	return model.Node{}, nil
}

func (n NodeService) CreateNode(ctx context.Context, newNode model.CreateNode) (model.Node, error) {

	return model.Node{}, nil
}
