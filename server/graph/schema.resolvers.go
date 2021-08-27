package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aligator/godrop/server/graph/dto"
	"github.com/aligator/godrop/server/graph/generated"
)

func (r *mutationResolver) CreateFileNode(ctx context.Context, input dto.CreateFileNode) (*dto.FileNode, error) {
	fileNode, err := r.FileNodeService.CreateFileNode(ctx, dto.CreateFileNodeFromDTO(input))

	if err != nil {
		return nil, err
	}

	res := dto.FileNodeFromModel(fileNode)
	return &res, nil
}

func (r *queryResolver) GetFileNode(ctx context.Context, path string) (*dto.FileNode, error) {
	fileNode, err := r.FileNodeService.GetFileNodeByPath(ctx, path)

	if err != nil {
		return nil, err
	}

	res := dto.FileNodeFromModel(fileNode)
	return &res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
