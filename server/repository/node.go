package repository

import (
	"context"

	"github.com/aligator/godrop/server/repository/model"
)

type Note interface {
	GetFileNodeByPath(ctx context.Context, path string) (model.FileNode, error)
	GetFileNodeById(ctx context.Context, id string) (model.FileNode, error)
	CreateFileNode(ctx context.Context, newFileNode model.CreateFileNode) (model.FileNode, error)
}
