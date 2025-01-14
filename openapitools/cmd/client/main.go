package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grigagod/go-openapi-example/openapitools/gen/client"
)

var (
	serverURL = "localhost:8080"
	secretKey = "qwerty"
	passenger = "Jan Kowalski"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := &oas.Configuration{
		Host:             "",
		Scheme:           "",
		DefaultHeader:    map[string]string{},
		UserAgent:        "",
		Debug:            false,
		Servers:          []oas.ServerConfiguration{},
		OperationServers: map[string]oas.ServerConfigurations{},
		HTTPClient:       http.DefaultClient,
	}
	cli := oas.NewAPIClient(cfg)

	var (
		originID  string
		destID    string
		originTZ  *time.Location
		tripID    string
		bookingID string
	)

	{
		req := cli.StationsAPI.
			GetStations(context.TODO()).
			Page(1).
			Limit(2).
			Country("PL")

		res, _, err := cli.StationsAPI.GetStationsExecute(req)
		if err != nil {
			panic(err)
		}

		origin, dest := res.Data[0], res.Data[1]
		originID, destID = origin.Id, dest.Id
		originTZ, _ = time.LoadLocation(origin.GetTimezone())
	}

	{
		req := cli.TripsAPI.GetTrips(ctx).
			Page(1).
			Limit(1).
			Origin(originID).
			Destination(destID).
			Date(time.Now().In(originTZ)).
			Dogs(true)
		res, _, err := cli.TripsAPI.GetTripsExecute(req)
		if err != nil {
			panic(err)
		}
		tripID = res.Data[0].GetId()
	}

	{
		req := cli.BookingsAPI.CreateBooking(ctx).
			Booking(oas.Booking{
				TripId:        oas.PtrString(tripID),
				PassengerName: oas.PtrString(passenger),
				HasDog:        oas.PtrBool(true),
			})
		res, _, err := cli.BookingsAPI.CreateBookingExecute(req)
		if err != nil {
			panic(err)
		}
		bookingID = res.GetId()
	}

	{
		req := cli.PaymentsAPI.CreateBookingPayment(ctx, bookingID).
			BookingPayment(oas.BookingPayment{
				Amount:   oas.PtrFloat32(15.51),
				Currency: oas.PtrString("PLN"),
				Source: &oas.BookingPaymentSource{
					Card: oas.NewCard(
						passenger,
						"4000006160000005",
						"777",
						2,
						2025,
						"PL",
					),
				},
			})
		res, _, err := cli.PaymentsAPI.CreateBookingPaymentExecute(req)
		if err != nil {
			panic(err)
		}

		if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
			panic(err)
		}
	}
}
