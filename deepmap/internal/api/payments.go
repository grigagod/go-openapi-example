package api

import (
	"context"

	"deepmap/gen/oas"
)

// Pay for a Booking
// (POST /bookings/{bookingId}/payment)
func (h Handler) CreateBookingPayment(ctx context.Context, request oas.CreateBookingPaymentRequestObject) (oas.CreateBookingPaymentResponseObject, error) {
	return oas.CreateBookingPayment200JSONResponse{}, nil
}
