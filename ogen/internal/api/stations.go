package api

import (
	"context"

	"ogen/gen/oas"
)

// GetStations implements get-stations operation.
//
// Returns a paginated and searchable list of all train stations.
//
// GET /stations
func (h Handler) GetStations(ctx context.Context, params oas.GetStationsParams) (oas.GetStationsRes, error) {
	return &oas.GetStationsOKHeaders{}, nil
}
