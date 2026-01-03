package app

import (
	"log"
	"net/http"
    "github.com/Tyulenb/kafka-order-manager/order-producer/internal/handler"
)

type App struct {
    Addr string
}

func NewApp(a string) *App {
    return &App{
        Addr: a,
    }
}

func (a *App) Run() error {
    mux := http.NewServeMux()

    producerHandler := handler.NewProducerHandler(nil)
    producerHandler.RegisterRoutes(mux)

    server := &http.Server{
        Addr: a.Addr,
        Handler: mux,
    }
    log.Printf("Server started on port :%s", a.Addr)
    return server.ListenAndServe()
}
