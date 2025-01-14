package main

import (
	"net/http"

	"github.com/grigagod/go-openapi-example/openapitools/internal/api"

	oas "github.com/grigagod/go-openapi-example/openapitools/gen/server/go"
)

func main() {
	h := api.Handler{}
	stationsCtl := oas.NewStationsAPIController(h)
	tripsCtl := oas.NewTripsAPIController(h)
	bookingsCtl := oas.NewBookingsAPIController(h)
	paymentsCtl := oas.NewPaymentsAPIController(h)

	mux := oas.NewRouter(
		stationsCtl,
		tripsCtl,
		bookingsCtl,
		paymentsCtl,
	)
	srv := http.Server{
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
