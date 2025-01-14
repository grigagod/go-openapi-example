package api

import (
	"context"

	"deepmap/gen/oas"
)

// List existing bookings
// (GET /bookings)
func (h Handler) GetBookings(ctx context.Context, request oas.GetBookingsRequestObject) (oas.GetBookingsResponseObject, error) {
	return oas.GetBookings200JSONResponse{}, nil
}

// Create a booking
// (POST /bookings)
func (h Handler) CreateBooking(ctx context.Context, request oas.CreateBookingRequestObject) (oas.CreateBookingResponseObject, error) {
	return oas.CreateBooking201JSONResponse{}, nil
}

// Delete a booking
// (DELETE /bookings/{bookingId})
func (h Handler) DeleteBooking(ctx context.Context, request oas.DeleteBookingRequestObject) (oas.DeleteBookingResponseObject, error) {
	return oas.DeleteBooking204Response{}, nil
}

// Get a booking
// (GET /bookings/{bookingId})
func (h Handler) GetBooking(ctx context.Context, request oas.GetBookingRequestObject) (oas.GetBookingResponseObject, error) {
	return oas.GetBooking200JSONResponse{}, nil
}
