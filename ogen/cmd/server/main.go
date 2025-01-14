package main

import (
	"errors"
	"flag"
	"net/http"

	"ogen/gen/oas"
	"ogen/internal/api"
	"ogen/internal/security"
)

var (
	addr   string
	secret string
)

func main() {
	flag.StringVar(&addr, "addr", ":9093", "")
	flag.StringVar(&secret, "secret", "qwerty", "")
	flag.Parse()

	oasSrv, err := oas.NewServer(
		api.Handler{},
		security.NewHandler(secret),
		oas.WithErrorHandler(api.ErrorHandler),
	)
	if err != nil {
		panic(err)
	}
	httpSrv := http.Server{
		Addr:    addr,
		Handler: oasSrv,
	}
	defer httpSrv.Close()
	if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
