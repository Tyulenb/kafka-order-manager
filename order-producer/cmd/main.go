package main

import (
	"log"

	"github.com/Tyulenb/kafka-order-manager/order-producer/internal/app"
)

func main() {
    app := app.NewApp(":8080")
    log.Println(app.Run())
}
