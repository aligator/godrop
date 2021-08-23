package repository

import (
	"context"

	"github.com/aligator/godrop/server/repository/model"
)

type Note interface {
	GetNodeByPath(ctx context.Context, path string) (model.Node, error)
	GetNodeById(ctx context.Context, id string) (model.Node, error)
	CreateNode(ctx context.Context, newNode model.CreateNode) (model.Node, error)
}
