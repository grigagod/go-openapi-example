package api

import (
	"context"
	"net/http"
	tt "traintravel"

	"ogen/gen/oas"
)

// GetStations implements get-stations operation.
//
// Returns a paginated and searchable list of all train stations.
//
// GET /stations
func (h Handler) GetStations(ctx context.Context, params oas.GetStationsParams) (oas.GetStationsRes, error) {
	if params.Country.Value != tt.Country {
		return &oas.BadRequestHeaders{
			Response: oas.Problem{
				Title:  oas.NewOptString("not supported country"),
				Status: oas.NewOptInt(http.StatusBadRequest),
			},
		}, nil
	}

	var resp oas.GetStationsOKHeaders
	for _, st := range []tt.Station{tt.Origin, tt.Destination} {
		resp.Response.Data = append(resp.Response.Data, oas.GetStationsOKDataItem{
			ID:          st.ID,
			Name:        st.Name,
			Address:     st.Address,
			CountryCode: tt.Country,
			Timezone:    oas.NewOptString(tt.Timezone.String()),
		})
	}
	return &resp, nil
}
