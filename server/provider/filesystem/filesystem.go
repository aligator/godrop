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

func (p *Provider) readNode(path string, withChildren bool) (model.Node, error) {
	result := model.Node{}

	file, err := p.FS.Open(path)
	if err != nil {
		return model.Node{}, checkpoint.From(err)
	}

	nodeStat, err := file.Stat()
	if err != nil {
		return model.Node{}, checkpoint.From(err)
	}

	result.IsFolder = nodeStat.IsDir()
	result.Name = file.Name()

	// For now just use the path as id.
	// Later we will cache all files in a db and can use that id.
	result.ID = path

	result.MimeType = mime.TypeByExtension(filepath.Ext(result.Name))

	if result.IsFolder && withChildren {
		children, err := file.Readdirnames(0)
		if err != nil {
			return model.Node{}, checkpoint.From(err)
		}

		for _, child := range children {
			// Populate the children without the children.
			childNode, err := p.readNode(filepath.Join(path, child), false)
			if err != nil {
				return model.Node{}, checkpoint.From(err)
			}

			result.Children = append(result.Children, childNode)
		}
	}

	return result, nil
}

func (p *Provider) GetNodeByPath(ctx context.Context, path string) (model.Node, error) {
	return p.readNode(path, true)
}

func (p *Provider) GetNodeById(ctx context.Context, id string) (model.Node, error) {
	// For now the id is also the path.
	return p.readNode(id, true)
}

func (p *Provider) CreateNode(ctx context.Context, newNode model.CreateNode) (model.Node, error) {
	panic("implement me")
}
