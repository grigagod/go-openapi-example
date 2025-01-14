package api

import (
	"context"

	"deepmap/gen/oas"
)

// Get a list of train stations
// (GET /stations)
func (h Handler) GetStations(ctx context.Context, request oas.GetStationsRequestObject) (oas.GetStationsResponseObject, error) {
	return oas.GetStations200JSONResponse{}, nil
}
