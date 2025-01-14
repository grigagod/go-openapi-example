package api

import (
	"context"
	"deepmap/gen/oas"
	"time"
	"traintravel"
	tt "traintravel"

	"github.com/google/uuid"
)

// Get available train trips
// (GET /trips)
func (h Handler) GetTrips(ctx context.Context, request oas.GetTripsRequestObject) (oas.GetTripsResponseObject, error) {
	var resp oas.GetTrips200JSONResponse

	if request.Params.Origin != traintravel.Origin.ID ||
		request.Params.Destination != traintravel.Destination.ID {
		return resp, nil
	}

	trips := []struct {
		ArrivalTime     *time.Time "json:\"arrival_time,omitempty\""
		BicyclesAllowed *bool      "json:\"bicycles_allowed,omitempty\""
		DepartureTime   *time.Time "json:\"departure_time,omitempty\""
		Destination     *string    "json:\"destination,omitempty\""
		DogsAllowed     *bool      "json:\"dogs_allowed,omitempty\""
		Id              *uuid.UUID "json:\"id,omitempty\""
		Operator        *string    "json:\"operator,omitempty\""
		Origin          *string    "json:\"origin,omitempty\""
		Price           *float32   "json:\"price,omitempty\""
		Self            *string    "json:\"self,omitempty\""
	}{
		{
			ArrivalTime:     &tt.NextTrip.ArrivalTime,
			BicyclesAllowed: &tt.NextTrip.BicyclesAllowed,
			DepartureTime:   &tt.NextTrip.DepartureTime,
			Destination:     toStrPtr(tt.NextTrip.Destination.ID.String()),
			DogsAllowed:     &tt.NextTrip.DogsAllowed,
			Id:              &tt.NextTrip.ID,
			Operator:        &tt.NextTrip.Operator,
			Origin:          toStrPtr(tt.NextTrip.Origin.ID.String()),
			Price:           &tt.NextTrip.Price,
		},
	}
	resp.Body.Data = &trips
	return resp, nil
}
