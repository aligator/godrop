package repository

import (
	"context"
	"errors"

	"github.com/aligator/godrop/server/repository/model"
)

var (
	ErrFileAlreadyExists = errors.New("file already exists")
)

type FileNode interface {
	GetFileNodeByPath(ctx context.Context, path string) (model.FileNode, error)
	GetFileNodeById(ctx context.Context, id string) (model.FileNode, error)

	// CreateFileNode creates the given file.
	// It may return ErrFileAlreadyExists.
	CreateFileNode(ctx context.Context, newFileNode model.CreateFileNode) (model.FileNode, error)

	SetState(ctx context.Context, id string, newState model.NodeState) error
}
