package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, msgDelay time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case <-time.After(msgDelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation done", msg)
				return
			case <-ctx.Done():
				out <- fmt.Sprintf("%v aborted", msg)
				return
			}
		}
	}()
	return out
}

func main() {
	ctx := context.Background()

	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := sampleOperation(ctx, "webserver", 200)
	microservice := sampleOperation(ctx, "microservice", 500)
	database := sampleOperation(ctx, "database", 900)

MainLoop:
	for {
		select {
		case val := <-webServer:
			fmt.Println(val)
		case val := <-microservice:
			fmt.Println(val)
			fmt.Println("cancel context")
			cancelCtx()
			break MainLoop
		case val := <-database:
			fmt.Println(val)
		}
	}
	fmt.Println(<-database)
}
