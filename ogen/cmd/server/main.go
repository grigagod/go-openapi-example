package main

import (
	"errors"
	"net/http"

	"ogen/gen/oas"
	"ogen/internal/api"
	"ogen/internal/security"
)

func main() {
	oasSrv, err := oas.NewServer(
		api.Handler{},
		security.Handler{SecretKey: "qwerty"},
		oas.WithErrorHandler(api.ErrorHandler),
	)
	if err != nil {
		panic(err)
	}
	httpSrv := http.Server{
		Handler: oasSrv,
	}
	defer httpSrv.Close()
	if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
