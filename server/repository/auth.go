package repository

import "context"

type Auth interface {
	Validate(ctx context.Context, bearerToken string) (bool, error)
}
