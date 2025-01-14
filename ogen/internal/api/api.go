package api

import (
	"context"
	"net/http"

	"ogen/gen/oas"
)

var (
	_ oas.Handler = Handler{}
)

type Handler struct {
	oas.UnimplementedHandler
}

func ErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {

}
