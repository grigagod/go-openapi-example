package client

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"time"

	"ogen/gen/oas"
	"ogen/internal/security"

	"github.com/google/uuid"
)

var (
	serverURL = "localhost:8080"
	secretKey = "qwerty"
	passenger = "Jan Kowalski"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cli, err := oas.NewClient(serverURL, security.Source{SecretKey: secretKey})
	if err != nil {
		panic(err)
	}

	var (
		originID  uuid.UUID
		destID    uuid.UUID
		originTZ  *time.Location
		tripID    uuid.UUID
		bookingID uuid.UUID
	)

	{
		r, ok := mustCall(cli.GetStations, ctx, oas.GetStationsParams{
			Page:    oas.NewOptInt(1),
			Limit:   oas.NewOptInt(2),
			Country: oas.NewOptString("PL"),
		}).(*oas.GetStationsOKHeaders)
		if !ok {
			panic(r)
		}
		origin, dest := r.Response.Data[0], r.Response.Data[1]
		originID, destID = origin.ID, dest.ID
		originTZ, _ = time.LoadLocation(origin.Timezone.Value)
	}

	{
		r, ok := mustCall(cli.GetTrips, ctx, oas.GetTripsParams{
			Page:        oas.NewOptInt(1),
			Limit:       oas.NewOptInt(1),
			Origin:      originID,
			Destination: destID,
			Date:        time.Now().In(originTZ),
			Dogs:        oas.NewOptBool(true),
		}).(*oas.GetTripsOKHeaders)
		if !ok {
			panic(r)
		}
		tripID = r.Response.Data[0].ID.Value
	}

	{
		r, ok := mustCall(cli.CreateBooking, ctx, &oas.Booking{
			TripID:        oas.NewOptUUID(tripID),
			PassengerName: oas.NewOptString(passenger),
			HasDog:        oas.NewOptBool(true),
		}).(*oas.CreateBookingCreated)
		if !ok {
			panic(r)
		}
		bookingID = r.ID.Value
	}

	{
		res, ok := mustCall2(cli.CreateBookingPayment, ctx,
			&oas.BookingPayment{
				Amount:   oas.NewOptFloat64(15.51),
				Currency: oas.NewOptBookingPaymentCurrency(oas.BookingPaymentCurrencyBam),
				Source: oas.NewOptBookingPaymentSource(
					oas.NewBookingPaymentSource0BookingPaymentSource(
						oas.BookingPaymentSource0{
							Name:           passenger,
							Number:         "4000006160000005",
							Cvc:            "777",
							ExpMonth:       2,
							ExpYear:        2025,
							AddressCountry: "PL",
						},
					),
				),
			}, oas.CreateBookingPaymentParams{
				BookingId: bookingID,
			}).(*oas.CreateBookingPaymentOKHeaders)
		if !ok {
			panic(res)
		}

		if err := json.NewEncoder(os.Stdout).Encode(res.Response); err != nil {
			panic(err)
		}
	}
}

func mustCall[In, Out any](inv func(context.Context, In) (Out, error), ctx context.Context, in In) Out {
	out, err := inv(ctx, in)
	if err != nil {
		panic(err)
	}
	return out
}

func mustCall2[In1, In2, Out any](inv func(context.Context, In1, In2) (Out, error), ctx context.Context, in1 In1, in2 In2) Out {
	out, err := inv(ctx, in1, in2)
	if err != nil {
		panic(err)
	}
	return out
}
