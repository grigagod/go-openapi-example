package security

import (
	"context"
	"time"

	"ogen/gen/oas"

	"github.com/golang-jwt/jwt"
)

var ss oas.SecuritySource = Source{}

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
	token := jwt.NewWithClaims(&jwt.SigningMethodHMAC{}, &Claims{
		Scopes: scopes,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour).Unix(),
			Id:        now.String(),
			IssuedAt:  now.Unix(),
		},
	})

	return oas.OAuth2{Token: token.Raw}, nil
}
