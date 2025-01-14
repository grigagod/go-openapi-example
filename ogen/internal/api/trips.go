package api

import (
	"context"
	"ogen/gen/oas"
)

// GetTrips implements get-trips operation.
//
// Returns a list of available train trips between the specified origin and destination stations on
// the given date, and allows for filtering by bicycle and dog allowances.
//
// GET /trips
func (h Handler) GetTrips(ctx context.Context, params oas.GetTripsParams) (oas.GetTripsRes, error) {
	return &oas.GetTripsOKHeaders{}, nil
}
