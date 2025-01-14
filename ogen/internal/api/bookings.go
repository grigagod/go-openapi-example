package api

import (
	"context"

	"ogen/gen/oas"
)

// CreateBooking implements create-booking operation.
//
// A booking is a temporary hold on a trip. It is not confirmed until the payment is processed.
//
// POST /bookings
func (h Handler) CreateBooking(ctx context.Context, req *oas.Booking) (oas.CreateBookingRes, error) {
	return &oas.CreateBookingCreated{}, nil
}

// DeleteBooking implements delete-booking operation.
//
// Deletes a booking, cancelling the hold on the trip.
//
// DELETE /bookings/{bookingId}
func (h Handler) DeleteBooking(ctx context.Context, params oas.DeleteBookingParams) (oas.DeleteBookingRes, error) {
	return &oas.DeleteBookingNoContent{}, nil
}

// GetBooking implements get-booking operation.
//
// Returns the details of a specific booking.
//
// GET /bookings/{bookingId}
func (h Handler) GetBooking(ctx context.Context, params oas.GetBookingParams) (oas.GetBookingRes, error) {
	return &oas.GetBookingOKHeaders{}, nil
}

// GetBookings implements get-bookings operation.
//
// Returns a list of all trip bookings by the authenticated user.
//
// GET /bookings
func (h Handler) GetBookings(ctx context.Context, params oas.GetBookingsParams) (oas.GetBookingsRes, error) {
	return &oas.GetBookingsOKHeaders{}, nil
}
