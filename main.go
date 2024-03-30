package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/mbhaskar98/orders-api/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	err := application.New().Start(ctx)

	if err != nil {
		fmt.Println("Error occured while starting server:", err)
	}
}
