package main

import (
	"context"
	"fmt"

	"github.com/sosamjayo/oauth/api/db"
)

func main() {
	fmt.Println("OAuth Server is running...")

	ctx := context.Background()
	_, err := db.NewRedisClient(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
}
