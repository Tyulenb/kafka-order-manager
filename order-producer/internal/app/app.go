package app

import (
	"log"
	"net/http"

	"github.com/Tyulenb/kafka-order-manager/order-producer/internal/handler"
	"github.com/Tyulenb/kafka-order-manager/order-producer/internal/service"
	"github.com/segmentio/kafka-go"
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

    w := &kafka.Writer {
        Addr: kafka.TCP("localhost:9092"),
        Topic: "topic-1",
        Balancer: &kafka.LeastBytes{},
        AllowAutoTopicCreation: true,
    }
    defer w.Close()

    producerService := service.NewProducerService(w)

    producerHandler := handler.NewProducerHandler(producerService)
    producerHandler.RegisterRoutes(mux)

    server := &http.Server{
        Addr: a.Addr,
        Handler: mux,
    }
    log.Printf("Server started on port :%s", a.Addr)
    return server.ListenAndServe()
}
