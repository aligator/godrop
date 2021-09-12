package filesystem

import (
	"context"
	"errors"
	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/repository"
	"github.com/aligator/godrop/server/repository/model"
	"github.com/spf13/afero"
	"mime"
	"path/filepath"
	"strings"
)

var (
	ErrNameNotAllowed = errors.New("the name is not allowed")
)

type NodeProvider struct {
	FS afero.Fs
}

func (p NodeProvider) readFileNode(path string, withChildren bool) (model.FileNode, error) {
	result := model.FileNode{}

	file, err := openById(p.FS, path)
	if err != nil {
		return model.FileNode{}, checkpoint.From(err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return model.FileNode{}, checkpoint.From(err)
	}

	result.IsFolder = fileStat.IsDir()
	result.Name = filepath.Base(file.Name())

	// For now just use the path as id.
	// Later we will cache all files in a db and can use that id.
	result.ID = path

	result.MimeType = mime.TypeByExtension(filepath.Ext(result.Name))

	if result.IsFolder && withChildren {
		children, err := file.Readdirnames(0)
		if err != nil {
			return model.FileNode{}, checkpoint.From(err)
		}

		for _, child := range children {
			// Populate the children without the children.
			childFileNode, err := p.readFileNode(filepath.Join(path, child), false)
			if err != nil {
				return model.FileNode{}, checkpoint.From(err)
			}

			result.Children = append(result.Children, childFileNode)
		}
	}

	result.Size = fileStat.Size()

	if strings.HasSuffix(result.Name, uploadSuffix) {
		result.State = model.NodeStateUpload
		result.Name = trimStateSuffix(result.Name)
	} else if strings.HasSuffix(result.Name, deleteSuffix) {
		return model.FileNode{}, checkpoint.From(repository.ErrFileNotFound)
	} else {
		result.State = model.NodeStateReady
	}

	return result, nil
}

func (p NodeProvider) GetByPath(_ context.Context, path string) (model.FileNode, error) {
	return p.readFileNode(path, true)
}

func (p NodeProvider) GetById(_ context.Context, id model.ID) (model.FileNode, error) {
	// For now the id is also the path.
	return p.readFileNode(id, true)
}

func (p NodeProvider) Create(ctx context.Context, newFileNode model.CreateFileNode) (model.FileNode, error) {
	// Check file name.
	if strings.HasSuffix(newFileNode.Name, stateSuffix) {
		return model.FileNode{}, checkpoint.From(ErrNameNotAllowed)
	}

	newFileId := filepath.Join(newFileNode.Path, newFileNode.Name)

	// Check if the file already exists.
	_, err := openById(p.FS, newFileId)
	if err != nil && errors.Is(err, afero.ErrFileNotFound) {
		return model.FileNode{}, checkpoint.Wrap(err, repository.ErrFileAlreadyExists)
	} else if err != nil {
		return model.FileNode{}, checkpoint.From(err)
	}

	if newFileNode.IsFolder {
		err := p.FS.Mkdir(newFileId, 0755)
		if err != nil {
			return model.FileNode{}, checkpoint.From(err)
		}
	} else {
		newFilePath := getStateFilename(newFileId, uploadSuffix)

		_, err := p.FS.Create(newFilePath)
		if errors.Is(err, afero.ErrFileExists) {
			return model.FileNode{}, checkpoint.Wrap(err, repository.ErrFileAlreadyExists)
		} else if err != nil {
			return model.FileNode{}, checkpoint.From(err)
		}
	}

	return p.GetByPath(ctx, newFileId)
}

func (p NodeProvider) SetState(ctx context.Context, id model.ID, newState model.NodeState) error {
	file, err := p.GetById(ctx, id)
	if err != nil {
		return checkpoint.From(err)
	}

	oldName := id
	newName := id

	if file.State == model.NodeStateUpload {
		oldName = getStateFilename(id, uploadSuffix)
	}

	if newState == model.NodeStateUpload {
		newName = getStateFilename(id, uploadSuffix)
	}

	return checkpoint.From(p.FS.Rename(oldName, newName))
}

func (p NodeProvider) DeleteById(_ context.Context, id model.ID) error {
	file, err := openById(p.FS, id)
	if err != nil {
		return checkpoint.From(err)
	}
	defer file.Close()

	err = p.FS.Rename(file.Name(), getStateFilename(file.Name(), deleteSuffix))
	return checkpoint.From(err)
}
