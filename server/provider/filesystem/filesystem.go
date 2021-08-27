package filesystem

import (
	"context"
	"mime"
	"os"
	"path/filepath"

	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/repository/model"
	"github.com/spf13/afero"
)

const defaultPath = "./files"

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
	if err != nil {
		return model.FileNode{}, checkpoint.From(err)
	}

	FileNodeStat, err := file.Stat()
	if err != nil {
		return model.FileNode{}, checkpoint.From(err)
	}

	result.IsFolder = FileNodeStat.IsDir()
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
	panic("implement me")
}
