package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type KafkaService interface {
    sendOrder(v any) error 
}

type ProducerHandler struct {
    kfk KafkaService
}

func NewProducerHandler(kfk KafkaService) *ProducerHandler {
    return &ProducerHandler{
        kfk: kfk,
    }
}

func (p *ProducerHandler) RegisterRoutes(router *http.ServeMux) {
    router.HandleFunc("POST /order", p.createOrder)
}

func (p *ProducerHandler) createOrder(w http.ResponseWriter, r *http.Request) {
    order := struct {
       Name string
       Amount int
    }{}
    if err := decodeJSON(r, &order); err != nil {
        http.Error(w, "Something went wrong!", http.StatusBadRequest)
        log.Printf("createOrder, err: %v\n", err)
        return
    }
    fmt.Println(order)
    p.kfk.sendOrder(order)
}

func decodeJSON(r *http.Request, v any) error {
    if r.Body == nil {
        return fmt.Errorf("The body is nil!")
    }
    return json.NewDecoder(r.Body).Decode(v)
}
