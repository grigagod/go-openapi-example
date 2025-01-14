package api

import (
	"context"
	"errors"
	"net/http"

	oas "github.com/grigagod/go-openapi-example/openapitools/gen/server/go"
)

// GetBookings - List existing bookings
func (h Handler) GetBookings(ctx context.Context, page int32, limit int32) (oas.ImplResponse, error) {
	// TODO - update GetBookings with the required logic for this service method.
	// Add api_bookings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetBookings200Response{}) or use other options such as http.Ok ...
	// return Response(200, GetBookings200Response{}), nil

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

	return oas.Response(http.StatusNotImplemented, nil), errors.New("GetBookings method not implemented")
}

// CreateBooking - Create a booking
func (h Handler) CreateBooking(ctx context.Context, booking oas.Booking) (oas.ImplResponse, error) {
	// TODO - update CreateBooking with the required logic for this service method.
	// Add api_bookings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, CreateBooking201Response{}) or use other options such as http.Ok ...
	// return Response(201, CreateBooking201Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(404, Problem{}) or use other options such as http.Ok ...
	// return Response(404, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(409, Problem{}) or use other options such as http.Ok ...
	// return Response(409, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(429, Problem{}) or use other options such as http.Ok ...
	// return Response(429, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(500, Problem{}) or use other options such as http.Ok ...
	// return Response(500, Problem{}), nil

	return oas.Response(http.StatusNotImplemented, nil), errors.New("CreateBooking method not implemented")
}

// GetBooking - Get a booking
func (h Handler) GetBooking(ctx context.Context, bookingId string) (oas.ImplResponse, error) {
	// TODO - update GetBooking with the required logic for this service method.
	// Add api_bookings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CreateBooking201Response{}) or use other options such as http.Ok ...
	// return Response(200, CreateBooking201Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(404, Problem{}) or use other options such as http.Ok ...
	// return Response(404, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(429, Problem{}) or use other options such as http.Ok ...
	// return Response(429, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(500, Problem{}) or use other options such as http.Ok ...
	// return Response(500, Problem{}), nil

	return oas.Response(http.StatusNotImplemented, nil), errors.New("GetBooking method not implemented")
}

// DeleteBooking - Delete a booking
func (h Handler) DeleteBooking(ctx context.Context, bookingId string) (oas.ImplResponse, error) {
	// TODO - update DeleteBooking with the required logic for this service method.
	// Add api_bookings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(404, Problem{}) or use other options such as http.Ok ...
	// return Response(404, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(429, Problem{}) or use other options such as http.Ok ...
	// return Response(429, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(500, Problem{}) or use other options such as http.Ok ...
	// return Response(500, Problem{}), nil

	return oas.Response(http.StatusNotImplemented, nil), errors.New("DeleteBooking method not implemented")
}
