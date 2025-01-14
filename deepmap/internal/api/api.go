package api

import (
	"deepmap/gen/oas"
)

var h oas.StrictServerInterface = Handler{}

type Handler struct{}

func toStrPtr(v string) *string {
	return &v
}

func badRequest(title string) oas.BadRequestApplicationProblemPlusJSONResponse {
	return oas.BadRequestApplicationProblemPlusJSONResponse{
		Body: oas.Problem{
			Title: &title,
		},
	}
}
