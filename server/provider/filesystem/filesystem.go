package filesystem

import (
	"context"
	"errors"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/repository"
	"github.com/aligator/godrop/server/repository/model"
	"github.com/spf13/afero"
)

const defaultPath = "./files"
const uploadSuffix = ".uploading"

var (
	ErrNameNotAllowed = errors.New("the name is not allowed")
)

type Provider struct {
	FS afero.Fs
}

func (p *Provider) Init() error {
	if p.FS == nil {

		if err := os.Mkdir(defaultPath, 0775); !os.IsExist(err) {
			return checkpoint.From(err)
		}

		p.FS = afero.NewBasePathFs(afero.NewOsFs(), defaultPath)
	}

	return nil
}

func (p *Provider) readFileNode(path string, withChildren bool) (model.FileNode, error) {
	result := model.FileNode{}

	file, err := p.FS.Open(path)
	if errors.Is(err, afero.ErrFileNotFound) {
		// Try the uploading file.
		return p.readFileNode(path+uploadSuffix, withChildren)
	} else if err != nil {
		return model.FileNode{}, checkpoint.From(err)
	}

	fileStat, err := file.Stat()
	if err != nil {
		return model.FileNode{}, checkpoint.From(err)
	}

	result.IsFolder = fileStat.IsDir()
	result.Name = filepath.Base(file.Name())

	// For now just use the path as id.
	// Later we will cache all files in a db and can use that id.
	result.ID = strings.TrimSuffix(file.Name(), uploadSuffix)

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
		result.Name = strings.TrimSuffix(result.Name, uploadSuffix)
	} else {
		result.State = model.NodeStateReady
	}

	return result, nil
}

func (p *Provider) GetFileNodeByPath(ctx context.Context, path string) (model.FileNode, error) {
	return p.readFileNode(path, true)
}

func (p *Provider) GetFileNodeById(ctx context.Context, id string) (model.FileNode, error) {
	// For now the id is also the path.
	return p.readFileNode(id, true)
}

func (p *Provider) CreateFileNode(ctx context.Context, newFileNode model.CreateFileNode) (model.FileNode, error) {
	// Check file name.
	if strings.HasSuffix(newFileNode.Name, uploadSuffix) {
		return model.FileNode{}, checkpoint.From(ErrNameNotAllowed)
	}

	newFileId := filepath.Join(newFileNode.Path, newFileNode.Name)

	if newFileNode.IsFolder {
		err := p.FS.Mkdir(newFileId, 0755)
		if err != nil {
			return model.FileNode{}, checkpoint.From(err)
		}
	} else {
		newFilePath := newFileId + uploadSuffix

		_, err := p.FS.Create(newFilePath)
		if errors.Is(err, afero.ErrFileExists) {
			return model.FileNode{}, checkpoint.Wrap(err, repository.ErrFileAlreadyExists)
		} else if err != nil {
			return model.FileNode{}, checkpoint.From(err)
		}
	}

	return p.GetFileNodeByPath(ctx, newFileId)
}

func (p *Provider) SetState(ctx context.Context, id string, newState model.NodeState) error {
	file, err := p.GetFileNodeById(ctx, id)
	if err != nil {
		return checkpoint.From(err)
	}

	oldName := id
	newName := id

	if file.State == model.NodeStateUpload {
		oldName = id + uploadSuffix
	}

	if newState == model.NodeStateUpload {
		newName = id + uploadSuffix
	}

	return checkpoint.From(p.FS.Rename(oldName, newName))
}

func (p *Provider) Save(ctx context.Context, id string, reader io.Reader) error {
	file, err := p.FS.OpenFile(id+uploadSuffix, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		return checkpoint.From(err)
	}

	_, err = io.Copy(file, reader)
	return checkpoint.From(err)
}

func (p *Provider) Read(ctx context.Context, id string, writer io.Writer) error {
	file, err := p.FS.Open(id)
	if err != nil {
		return checkpoint.From(err)
	}

	_, err = io.Copy(writer, file)
	return checkpoint.From(err)
}
