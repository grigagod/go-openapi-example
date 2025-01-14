package api

import (
	"context"
	"deepmap/gen/oas"
)

// Get available train trips
// (GET /trips)
func (h Handler) GetTrips(ctx context.Context, request oas.GetTripsRequestObject) (oas.GetTripsResponseObject, error) {
	return oas.GetTrips200JSONResponse{}, nil
}
