package api

import (
	"context"
	"ogen/gen/oas"
	tt "traintravel"
)

// GetTrips implements get-trips operation.
//
// Returns a list of available train trips between the specified origin and destination stations on
// the given date, and allows for filtering by bicycle and dog allowances.
//
// GET /trips
func (h Handler) GetTrips(ctx context.Context, params oas.GetTripsParams) (oas.GetTripsRes, error) {
	var resp oas.GetTripsOKHeaders
	if params.Origin != tt.Origin.ID ||
		params.Destination != tt.Destination.ID {
		return &resp, nil
	}

	trip := tt.NextTrip
	resp.Response.Data = []oas.GetTripsOKDataItem{
		{
			ID:              oas.NewOptUUID(trip.ID),
			Origin:          oas.NewOptString(trip.Origin.ID.String()),
			Destination:     oas.NewOptString(trip.Origin.ID.String()),
			DepartureTime:   oas.NewOptDateTime(trip.DepartureTime),
			ArrivalTime:     oas.NewOptDateTime(trip.ArrivalTime),
			Operator:        oas.NewOptString(trip.Operator),
			Price:           oas.NewOptFloat64(float64(trip.Price)),
			BicyclesAllowed: oas.NewOptBool(trip.BicyclesAllowed),
			DogsAllowed:     oas.NewOptBool(trip.DogsAllowed),
		},
	}
	return &resp, nil
}
