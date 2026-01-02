package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic: "topic-1",
        Partition: 0,
        MaxBytes: 10e6,
    })
    defer r.Close()

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            break
        }
        fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
    }
}
