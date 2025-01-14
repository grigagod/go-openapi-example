package api

import (
	"context"
	"net/http"
	tt "traintravel"

	"ogen/gen/oas"
)

// CreateBookingPayment implements create-booking-payment operation.
//
// A payment is an attempt to pay for the booking, which will confirm the booking for the user and
// enable them to get their tickets.
//
// POST /bookings/{bookingId}/payment
func (h Handler) CreateBookingPayment(ctx context.Context, req *oas.BookingPayment, params oas.CreateBookingPaymentParams) (oas.CreateBookingPaymentRes, error) {
	cardSrc, ok := req.Source.Value.GetBookingPaymentSource0()
	if !ok {
		return &oas.BadRequestHeaders{
			Response: oas.Problem{
				Title:  oas.NewOptString("only card payments supported"),
				Status: oas.NewOptInt(http.StatusBadRequest),
			},
		}, nil
	}
	booking, ok := tt.GetBooking(params.BookingId)
	if !ok {
		return &oas.BadRequestHeaders{
			Response: oas.Problem{
				Title:  oas.NewOptString("invalid booking_id"),
				Status: oas.NewOptInt(http.StatusBadRequest),
			},
		}, nil
	}
	payment := tt.NewPayment(
		booking.ID,
		float32(req.Amount.Value),
		string(req.Currency.Value),
		tt.Card{
			Name:     cardSrc.Name,
			Number:   cardSrc.Name,
			CVC:      cardSrc.Cvc,
			ExpMonth: cardSrc.ExpMonth,
			ExpYear:  cardSrc.ExpYear,
		},
	)
	return &oas.CreateBookingPaymentOKHeaders{
		Response: oas.CreateBookingPaymentOK{
			ID:     oas.NewOptUUID(payment.ID),
			Amount: oas.NewOptFloat64(float64(payment.Amount)),
			Currency: oas.NewOptCreateBookingPaymentOKCurrency(
				oas.CreateBookingPaymentOKCurrency(payment.Currency),
			),
			Source: oas.NewOptCreateBookingPaymentOKSource(
				oas.NewCreateBookingPaymentOKSource0CreateBookingPaymentOKSource(
					oas.CreateBookingPaymentOKSource0(cardSrc),
				),
			),
			Status: oas.NewOptCreateBookingPaymentOKStatus(
				oas.CreateBookingPaymentOKStatus(payment.Status),
			),
		},
	}, nil
}
