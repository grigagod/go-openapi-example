package main

import (
	"context"
	"deepmap/gen/oas"
	"deepmap/internal/api"
	"flag"
	"net/http"
	"strings"
	tt "traintravel"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"

	"github.com/oapi-codegen/nethttp-middleware"
)

var (
	addr   string
	secret string
)

func main() {
	flag.StringVar(&addr, "addr", ":9093", "")
	flag.StringVar(&secret, "secret", "qwerty", "")
	flag.Parse()

	doc, err := openapi3.NewLoader().LoadFromFile("openapi.yaml")
	if err != nil {
		panic(err)
	}
	doc.Servers = nil // disable Host header validation

	h := oas.HandlerWithOptions(
		oas.NewStrictHandler(api.Handler{}, nil),
		oas.StdHTTPServerOptions{
			BaseRouter: http.NewServeMux(),
			Middlewares: []oas.MiddlewareFunc{
				nethttpmiddleware.OapiRequestValidatorWithOptions(doc,
					&nethttpmiddleware.Options{
						Options: openapi3filter.Options{
							AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
								header := ai.RequestValidationInput.Request.Header.Get("Authorization")
								rawToken := strings.TrimPrefix(header, "Bearer ")
								return tt.HandleOAuth(secret, rawToken, ai.Scopes)
							},
						},
					}),
			},
		},
	)

	httpSrv := http.Server{
		Addr:    addr,
		Handler: h,
	}
	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err)
	}
}
