package main

import (
	"net"
	"net/http"
	"os"

	"github.com/najibulloShapoatov/http/cmd/app"
	"github.com/najibulloShapoatov/http/pkg/banners"
)

func main() {
	//обьявляем порт и хост
	host := "0.0.0.0"
	port := "9999"

	//вызываем фукцию execute если получили ошибку то закрываем программу с ошибкой
	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(h, p string) error {
	//создаём новый мукс для хандлеров
	mux := http.NewServeMux()
	//создаём новый сервис
	bnrSvc := banners.NewService()

	//создаём новый сервер с сервисами
	sr := app.NewServer(mux, bnrSvc)

	//инициализируем сервер и регистрируем новый роутеры
	sr.Init()

	//создаём новый HTTP server
	srv := &http.Server{
		Addr:    net.JoinHostPort(h, p),
		Handler: sr,
	}
	//запускаем сервер и если получим ошибку то его вернем в резултать
	return srv.ListenAndServe()
}
