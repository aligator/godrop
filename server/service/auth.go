package service

import (
	"context"

	"github.com/aligator/godrop/server/log"
	"github.com/aligator/godrop/server/provider"
)

type AuthService struct {
	Logger log.GoDropLogger
	Repos  *provider.Repositories
}

func (f AuthService) Validate(ctx context.Context, bearerToken string) (bool, error) {
	return f.Repos.Auth.Validate(ctx, bearerToken)
}
