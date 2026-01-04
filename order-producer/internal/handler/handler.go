package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type KafkaService interface {
    SendOrder(v []byte) error 
}

type ProducerHandler struct {
    service KafkaService
}

func NewProducerHandler(kfk KafkaService) *ProducerHandler {
    return &ProducerHandler{
        service: kfk,
    }
}

func (p *ProducerHandler) RegisterRoutes(router *http.ServeMux) {
    router.HandleFunc("POST /order", p.createOrder)
}

func (p *ProducerHandler) createOrder(w http.ResponseWriter, r *http.Request) {
    order, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Something went wrong!", http.StatusBadRequest)
        log.Printf("createOrder -> ReadAll, err: %v\n", err)
        return
    }
    fmt.Println(string(order))
    if err := p.service.SendOrder(order); err != nil {
        http.Error(w, "Something went wrong!", http.StatusBadRequest)
        log.Printf("createOrder -> SendOrder, err: %v\n", err)
        return
    }
}
