package api

import (
	"context"
	"time"

	oas "github.com/grigagod/go-openapi-example/openapitools/gen/server/go"
)

func (h Handler) GetTrips(
	ctx context.Context,
	origin string,
	destination string,
	date time.Time,
	page int32,
	limit int32,
	bicycles bool,
	dogs bool,
) (oas.ImplResponse, error) {
	return oas.Response(200, oas.GetTrips200Response{
		Data:  []oas.GetTrips200ResponseAllOfDataInner{},
		Links: oas.GetStations200ResponseAllOfLinks{},
	}), nil
}
