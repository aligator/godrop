package service

import (
	"context"
	"errors"
	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/log"
	"github.com/aligator/godrop/server/repository/model"
	"io"

	"github.com/aligator/godrop/server/provider"
)

var (
	ErrFileNotInUploadState = errors.New("the file is not in upload state")
)

type FileService struct {
	Logger log.GoDropLogger
	Repos  *provider.Repositories
}

func (f FileService) Upload(ctx context.Context, id string, reader io.Reader) error {
	node, err := f.Repos.Node.GetByPath(ctx, id)
	if err != nil {
		return checkpoint.From(err)
	}

	if node.State != model.NodeStateUpload {
		return checkpoint.From(ErrFileNotInUploadState)
	}

	err = f.Repos.File.Save(ctx, id, reader)
	if err != nil {
		return checkpoint.From(err)
	}

	// Set to ready.
	return checkpoint.From(f.Repos.Node.SetState(ctx, id, model.NodeStateReady))
}

func (f FileService) Download(ctx context.Context, id string, writer io.Writer) error {
	return f.Repos.File.Read(ctx, id, writer)
}
