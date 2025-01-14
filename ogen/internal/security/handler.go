package security

import (
	"context"
	"fmt"
	"ogen/gen/oas"
	"slices"

	"github.com/golang-jwt/jwt"
)

var sh oas.SecurityHandler = Handler{}

type Handler struct {
	SecretKey string
}

type Claims struct {
	Scopes []string
	jwt.StandardClaims
}

func (h Handler) HandleOAuth2(ctx context.Context, _ oas.OperationName, t oas.OAuth2) (context.Context, error) {
	token, err := jwt.ParseWithClaims(t.Token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(h.SecretKey), nil
	})
	if err != nil {
		return ctx, fmt.Errorf("can't parse token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return ctx, fmt.Errorf("invalid token")
	}

	for _, scope := range t.Scopes {
		if !slices.Contains(claims.Scopes, scope) {
			return ctx, fmt.Errorf("missing required scope %s", scope)
		}
	}
	return ctx, nil
}
