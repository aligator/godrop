package filesystem

import (
	"errors"
	"github.com/aligator/checkpoint"
	"github.com/spf13/afero"
)

var (
	ErrFilesystemNil = errors.New("the given filesystem is nil")
)

func New(filesystem afero.Fs) (NodeProvider, FileProvider, error) {
	if filesystem == nil {
		return NodeProvider{}, FileProvider{}, checkpoint.From(ErrFilesystemNil)
	}

	nodeProvider := NodeProvider{FS: filesystem}
	fileProvider := FileProvider{FS: filesystem}

	nodeProvider.FS = filesystem
	fileProvider.FS = filesystem

	return nodeProvider, fileProvider, nil
}
