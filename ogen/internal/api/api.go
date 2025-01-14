package api

import (
	"context"
	"net/http"

	"github.com/go-faster/jx"
	"github.com/ogen-go/ogen/ogenerrors"

	"ogen/gen/oas"
)

var (
	_ oas.Handler = Handler{}
)

type Handler struct {
	oas.UnimplementedHandler
}

func ErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	code := ogenerrors.ErrorCode(err)
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(code)

	e := jx.GetEncoder()
	problem := oas.Problem{
		Title:  oas.NewOptString(e.String()),
		Status: oas.NewOptInt(code),
	}

	problem.Encode(e)
	_, _ = w.Write(e.Bytes())
}
