package repository

import (
	"context"
	"errors"

	"github.com/aligator/godrop/server/repository/model"
)

var (
	ErrFileAlreadyExists = errors.New("file already exists")
	ErrFileNotFound      = errors.New("file not found")
)

type FileNode interface {
	// GetByPath reads the file node at the given path.
	// It may return ErrFileNotFound.
	GetByPath(ctx context.Context, path string) (model.FileNode, error)

	// GetById reads the file node with the given ID.
	// It may return ErrFileNotFound.
	GetById(ctx context.Context, id model.ID) (model.FileNode, error)

	// Create creates the given file.
	// It may return ErrFileAlreadyExists.
	Create(ctx context.Context, newFileNode model.CreateFileNode) (model.FileNode, error)

	// DeleteById deletes the node with the given ID.
	// It may return ErrFileNotFound.
	DeleteById(ctx context.Context, id model.ID) error

	// SetState changes the state of the node with the given ID.
	// It may return ErrFileNotFound.
	SetState(ctx context.Context, id model.ID, newState model.NodeState) error
}
