package api

import (
	"context"
	"errors"
	"net/http"

	oas "github.com/grigagod/go-openapi-example/openapitools/gen/server/go"
)

func (h Handler) CreateBookingPayment(ctx context.Context, bookingId string, bookingPayment oas.BookingPayment) (oas.ImplResponse, error) {
	// TODO - update CreateBookingPayment with the required logic for this service method.
	// Add api_payments_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CreateBookingPayment200Response{}) or use other options such as http.Ok ...
	// return Response(200, CreateBookingPayment200Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(429, Problem{}) or use other options such as http.Ok ...
	// return Response(429, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(500, Problem{}) or use other options such as http.Ok ...
	// return Response(500, Problem{}), nil

	return oas.Response(http.StatusNotImplemented, nil), errors.New("CreateBookingPayment method not implemented")
}
