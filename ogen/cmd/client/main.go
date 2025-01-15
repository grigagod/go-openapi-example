package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"time"

	"ogen/gen/oas"
	"ogen/internal/security"

	"github.com/go-faster/jx"
	"github.com/google/uuid"
)

var (
	url       = "http://127.0.0.1:9093"
	secret    = "qwerty"
	passenger = "Jan Kowalski"
	country   = "PL"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cli, err := oas.NewClient(url, security.Source{SecretKey: secret})
	if err != nil {
		panic(err)
	}

	var (
		originID  uuid.UUID
		destID    uuid.UUID
		originTZ  *time.Location = time.UTC
		tripID    uuid.UUID
		bookingID uuid.UUID
	)

	{
		res, err := cli.GetStations(ctx, oas.GetStationsParams{
			Page:    oas.NewOptInt(1),
			Limit:   oas.NewOptInt(2),
			Country: oas.NewOptString(country),
		})
		if err != nil {
			panic(err)
		}
		switch r := res.(type) {
		case *oas.GetStationsOKHeaders:
			origin, dest := r.Response.Data[0], r.Response.Data[1]
			originID, destID = origin.ID, dest.ID
			originTZ, _ = time.LoadLocation(origin.Timezone.Value)
		default:
			if err := json.NewEncoder(os.Stderr).Encode(r); err != nil {
				panic(err)
			}
		}
	}

	{
		res, err := cli.GetTrips(ctx, oas.GetTripsParams{
			Page:        oas.NewOptInt(1),
			Limit:       oas.NewOptInt(1),
			Origin:      originID,
			Destination: destID,
			Date:        time.Now().In(originTZ),
			Dogs:        oas.NewOptBool(true),
		})
		if err != nil {
			panic(err)
		}
		switch r := res.(type) {
		case *oas.GetTripsOKHeaders:
			if len(r.Response.Data) != 0 {
				tripID = r.Response.Data[0].ID.Value
			}
		default:
			if err := json.NewEncoder(os.Stderr).Encode(r); err != nil {
				panic(err)
			}
		}
	}

	{
		res, err := cli.CreateBooking(ctx, &oas.Booking{
			TripID:        oas.NewOptUUID(tripID),
			PassengerName: oas.NewOptString(passenger),
			HasDog:        oas.NewOptBool(true),
		})
		if err != nil {
			panic(err)
		}
		switch r := res.(type) {
		case *oas.CreateBookingCreated:
			bookingID = r.ID.Value
		default:
			if err := json.NewEncoder(os.Stderr).Encode(r); err != nil {
				panic(err)
			}
		}

	}

	{
		res, err := cli.CreateBookingPayment(ctx,
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
							AddressCountry: country,
						},
					),
				),
			}, oas.CreateBookingPaymentParams{
				BookingId: bookingID,
			})
		if err != nil {
			panic(err)
		}

		switch r := res.(type) {
		case *oas.CreateBookingPaymentOKHeaders:
			en := jx.GetEncoder()
			r.Response.Encode(en)
			fmt.Fprint(os.Stdout, en)
			jx.PutEncoder(en)
		default:
			panic(r)
		}
	}
}
