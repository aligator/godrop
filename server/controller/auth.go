package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/aligator/godrop/server/log"
	"github.com/aligator/godrop/server/service"
)

var (
	ErrInvalidAuthorizationHeader = errors.New("invalid Authorization header")
)

type AuthController struct {
	Logger      log.GoDropLogger
	AuthService *service.AuthService
}

func (ac AuthController) JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		const prefix = "Bearer "
		// Read the Bearer token.
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		if !strings.HasPrefix(authHeader, prefix) {
			http.Error(w, ErrInvalidAuthorizationHeader.Error(), http.StatusBadRequest)
			return
		}

		token := strings.TrimPrefix(authHeader, prefix)

		authorized, err := ac.AuthService.Validate(ctx, token)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			ac.Logger.Error(err)
			return
		}

		if !authorized {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
