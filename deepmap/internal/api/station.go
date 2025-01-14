package api

import (
	"context"
	tt "traintravel"

	"deepmap/gen/oas"
)

// Get a list of train stations
// (GET /stations)
func (h Handler) GetStations(ctx context.Context, request oas.GetStationsRequestObject) (oas.GetStationsResponseObject, error) {
	var resp oas.GetStations200JSONResponse
	if *request.Params.Country != tt.Country {
		return &oas.GetStations400ApplicationProblemPlusJSONResponse{
			badRequest("not supported country"),
		}, nil
	}
	stations := make([]oas.Station, 0, 1)
	for _, station := range []tt.Station{tt.Origin, tt.Destination} {
		stations = append(stations, oas.Station{
			Address:     station.Address,
			CountryCode: tt.Country,
			Id:          station.ID,
			Name:        station.Name,
			Timezone:    toStrPtr(tt.Timezone.String()),
		})
	}
	resp.Body.Data = &stations
	return resp, nil
}
