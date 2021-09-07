package repository

import (
	"context"
	"io"
)

type File interface {
	Save(ctx context.Context, id string, reader io.Reader) error
	Read(ctx context.Context, id string, writer io.Writer) error
}
