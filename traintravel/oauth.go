package traintravel

import (
	"fmt"
	"slices"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Scopes []string
	jwt.RegisteredClaims
}

func HandleOAuth(secret, raw string, scopes []string) error {
	token, err := jwt.ParseWithClaims(raw, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return fmt.Errorf("can't parse token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token")
	}

	for _, scope := range scopes {
		if !slices.Contains(claims.Scopes, scope) {
			return fmt.Errorf("missing scope %s", scope)
		}
	}
	return nil
}
