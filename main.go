package main

import (
	"context"
	"fmt"

	"github.com/mbhaskar98/orders-api/application"
)

func main() {
	err := application.New().Start(context.TODO())

	if err != nil {
		fmt.Println("Error occured while starting server:", err)
	}
}
