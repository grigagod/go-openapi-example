package api

import (
	"context"

	oas "github.com/grigagod/go-openapi-example/openapitools/gen/server/go"
)

// GetStations - Get a list of train stations
func (s Handler) GetStations(
	ctx context.Context,
	page int32,
	limit int32,
	coordinates string,
	search string,
	country string,
) (oas.ImplResponse, error) {
	return oas.Response(200, oas.GetStations200Response{
		Data:  []oas.Station{},
		Links: oas.GetStations200ResponseAllOfLinks{},
	}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(429, Problem{}) or use other options such as http.Ok ...
	// return Response(429, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(500, Problem{}) or use other options such as http.Ok ...
	// return Response(500, Problem{}), nil
}
