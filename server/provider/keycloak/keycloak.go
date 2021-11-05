package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v8"
)

// Config needed for keycloak.
type Config struct {
	Host     string
	Realm    string
	ClientID string
}

// keycloak for authentication support by keycloak.
// Create new instances using New.
type keycloak struct {
	client gocloak.GoCloak
	realm  string
}

// New creates a keycloak repository based on the given config.
func New(config Config) keycloak {
	return keycloak{
		client: gocloak.NewClient(config.Host),
		realm:  config.Realm,
	}
}

// Validate the given token. If the token could not be parsed, it returns false and the error.
// If it could validate the token, it returns if the token is valid.
func (a keycloak) Validate(bearerToken string) (bool, error) {
	ctx := context.Background()

	// For now it should be enough to just check the token offline -> less network traffic.
	jwt, _, err := a.client.DecodeAccessToken(ctx, bearerToken, a.realm, "account")

	if err != nil {
		return false, err
	}

	return jwt.Valid, nil
}
