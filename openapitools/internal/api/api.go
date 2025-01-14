package api

import (
	oas "github.com/grigagod/go-openapi-example/openapitools/gen/server/go"
)

var (
	_ oas.StationsAPIServicer = Handler{}
	_ oas.TripsAPIServicer    = Handler{}
	_ oas.BookingsAPIServicer = Handler{}
	_ oas.PaymentsAPIServicer = Handler{}
)

type Handler struct{}
