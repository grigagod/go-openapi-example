package security

import (
	"context"
	"fmt"
	"time"

	"ogen/gen/oas"

	"github.com/golang-jwt/jwt/v5"
)

var _ oas.SecuritySource = Source{}

type Source struct {
	SecretKey string
}

func (s Source) OAuth2(ctx context.Context, operationName oas.OperationName) (oas.OAuth2, error) {
	var scopes []string
	switch operationName {
	case oas.GetBookingOperation, oas.GetBookingsOperation, oas.GetStationsOperation, oas.GetTripsOperation:
		scopes = []string{"read"}
	case oas.CreateBookingOperation, oas.CreateBookingPaymentOperation, oas.DeleteBookingOperation:
		scopes = []string{"write"}
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Scopes: scopes,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	ss, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return oas.OAuth2{}, fmt.Errorf("cant sign token: %w", err)
	}

	return oas.OAuth2{Token: ss, Scopes: scopes}, nil
}
