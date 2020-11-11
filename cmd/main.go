package main

import (
	"net"
	"net/http"
	"os"

	"github.com/najibulloShapoatov/http/cmd/app"
	"github.com/najibulloShapoatov/http/pkg/banners"
)

func main() {
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(h, p string) error {
	mux := http.NewServeMux()
	bnrSvc := banners.NewService()

	sr := app.NewServer(mux, bnrSvc)
	sr.Init()

	srv := &http.Server{
		Addr:    net.JoinHostPort(h, p),
		Handler: sr,
	}
	return srv.ListenAndServe()
}
