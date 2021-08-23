package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aligator/godrop/server/graph/dto"
	"github.com/aligator/godrop/server/graph/generated"
)

func (r *mutationResolver) CreateNode(ctx context.Context, input dto.CreateNode) (*dto.Node, error) {
	node, err := r.NodeService.CreateNode(ctx, dto.CreateNodeFromDTO(input))

	if err != nil {
		return nil, err
	}

	res := dto.NodeFromModel(node)
	return &res, nil
}

func (r *queryResolver) GetNode(ctx context.Context, path string) (*dto.Node, error) {
	node, err := r.NodeService.GetNodeByPath(ctx, path)

	if err != nil {
		return nil, err
	}

	res := dto.NodeFromModel(node)
	return &res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
