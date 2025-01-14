package api

import (
	"context"
	tt "traintravel"

	"deepmap/gen/oas"
)

// Pay for a Booking
// (POST /bookings/{bookingId}/payment)
func (h Handler) CreateBookingPayment(ctx context.Context, request oas.CreateBookingPaymentRequestObject) (oas.CreateBookingPaymentResponseObject, error) {
	cardSrc, err := request.Body.Source.AsBookingPaymentSource0()
	if err != nil {
		return &oas.CreateBookingPayment400ApplicationProblemPlusJSONResponse{
			badRequest("only card payments supported"),
		}, nil
	}

	booking, ok := tt.GetBooking(request.BookingId)
	if !ok {
		return &oas.CreateBookingPayment400ApplicationProblemPlusJSONResponse{
			badRequest("invalid booking_id"),
		}, nil
	}

	payment := tt.NewPayment(
		booking.ID,
		*request.Body.Amount,
		string(*request.Body.Currency),
		tt.Card{
			Name:     cardSrc.Name,
			Number:   cardSrc.Number,
			CVC:      *cardSrc.Cvc,
			ExpMonth: cardSrc.ExpMonth,
			ExpYear:  cardSrc.ExpYear,
		},
	)
	return oas.CreateBookingPayment200JSONResponse{
		Body: oas.BookingPayment{
			Amount:   request.Body.Amount,
			Currency: request.Body.Currency,
			Id:       &payment.ID,
			Source:   request.Body.Source,
			Status:   (*oas.BookingPaymentStatus)(&payment.Status),
		},
	}, nil
}
