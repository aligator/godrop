package service

import (
	"context"
	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/repository/model"
	"io"

	"github.com/aligator/godrop/server/provider"
)

type FileService struct {
	Repos *provider.Repositories
}

func (f FileService) Upload(ctx context.Context, id string, reader io.Reader) error {
	err := f.Repos.File.Save(ctx, id, reader)
	if err != nil {
		return checkpoint.From(err)
	}

	// Set to ready.
	return checkpoint.From(f.Repos.Node.SetState(ctx, id, model.NodeStateReady))
}

func (f FileService) Download(ctx context.Context, id string, writer io.Writer) error {
	return f.Repos.File.Read(ctx, id, writer)
}
