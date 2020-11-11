package main

import (
	"github.com/najibulloShapoatov/http/pkg/banners"
	"github.com/najibulloShapoatov/http/cmd/app"
	"net/http"
	"os"
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

 srv := &http.Server{
	 Addr: net.JoinHostPort(h, p),
	 Handler: sr,
 }
 return srv.ListenAndServe()
}
