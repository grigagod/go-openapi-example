package security

import (
	"context"
	"ogen/gen/oas"
	tt "traintravel"

	"github.com/golang-jwt/jwt/v5"
)

var sh oas.SecurityHandler = Handler{}

type Handler struct {
	secret string
}

func NewHandler(secret string) Handler {
	return Handler{
		secret: secret,
	}
}

type Claims struct {
	Scopes []string
	jwt.RegisteredClaims
}

func (h Handler) HandleOAuth2(ctx context.Context, _ oas.OperationName, t oas.OAuth2) (context.Context, error) {
	return ctx, tt.HandleOAuth(h.secret, t.Token, t.Scopes)
}
