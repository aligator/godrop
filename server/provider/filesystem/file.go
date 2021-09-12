package filesystem

import (
	"context"
	"errors"
	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/repository/model"
	"github.com/spf13/afero"
	"io"
	"os"
)

var (
	ErrFileIsNotInDeleteState = errors.New("file is not in delete state")
)

type FileProvider struct {
	FS afero.Fs
}

func (p FileProvider) Save(_ context.Context, id model.ID, reader io.Reader) error {
	file, err := openFileById(p.FS, id, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		return checkpoint.From(err)
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	return checkpoint.From(err)
}

func (p FileProvider) Read(_ context.Context, id model.ID, writer io.Writer) error {
	// Here we only want to read if it is ready -> no need for the openById method.
	file, err := p.FS.Open(id)
	if err != nil {
		return checkpoint.From(err)
	}
	defer file.Close()

	_, err = io.Copy(writer, file)
	return checkpoint.From(err)
}

func (p FileProvider) Delete(_ context.Context, id string) error {
	file, err := openById(p.FS, id)
	if err != nil {
		return checkpoint.From(err)
	}
	defer file.Close()

	if getStateOf(file.Name()) != deleteSuffix {
		return checkpoint.From(ErrFileIsNotInDeleteState)
	}

	return checkpoint.From(p.FS.RemoveAll(file.Name()))
}
