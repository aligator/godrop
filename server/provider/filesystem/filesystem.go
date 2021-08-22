package filesystem

import (
	"context"

	"github.com/aligator/godrop/server/repository/model"
	"github.com/spf13/afero"
)

type Provider struct {
	FS afero.Fs
}

func (p *Provider) Init() error {
	if p.FS == nil {
		var err error
		p.FS = afero.NewBasePathFs(afero.NewOsFs(), ".")
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Provider) GetNode(ctx context.Context, path string) (model.Node, error) {
	panic("implement me")
}

func (p *Provider) CreateNode(ctx context.Context, newNode model.CreateNode) (model.Node, error) {
	panic("implement me")
}
