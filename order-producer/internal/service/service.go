package service 

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
    w := kafka.Writer {
        Addr: kafka.TCP("localhost:9092"),
        Topic: "topic-1",
        Balancer: &kafka.LeastBytes{},
        AllowAutoTopicCreation: true,
    }
    defer w.Close()

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    err := w.WriteMessages(ctx,
        kafka.Message {
            Key: []byte("Key-A"),
            Value: []byte("My first kafka message"),
        },
        kafka.Message {
            Key: []byte("Key-B"),
            Value: []byte("My second kafka message"),
        },
        kafka.Message {
            Key: []byte("Key-C"),
            Value: []byte("Another one kafka message..."),
        },
    )
    if err != nil {
        log.Fatal("Error to write messages!", err)
    }
}
