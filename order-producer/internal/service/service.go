package service 

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type ProducerService struct {
    writer *kafka.Writer
}

func NewProducerService(writer *kafka.Writer) *ProducerService {
    return &ProducerService{
        writer: writer,
    }
}

func (ps *ProducerService) SendOrder(order []byte) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()
    err := ps.writer.WriteMessages(ctx,
        kafka.Message{
            Key: []byte("new_order"), 
            Value: order,
        },
    )
    return err
}
