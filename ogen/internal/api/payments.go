package api

import (
	"context"

	"ogen/gen/oas"
)

// CreateBookingPayment implements create-booking-payment operation.
//
// A payment is an attempt to pay for the booking, which will confirm the booking for the user and
// enable them to get their tickets.
//
// POST /bookings/{bookingId}/payment
func (h Handler) CreateBookingPayment(ctx context.Context, req *oas.BookingPayment, params oas.CreateBookingPaymentParams) (oas.CreateBookingPaymentRes, error) {
	return &oas.CreateBookingPaymentOKHeaders{}, nil
}
