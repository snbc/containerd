package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	absTimeout := time.Now().Add(30 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), absTimeout)

	defer cancel()

	go func(ctx context.Context) {
		for {
			time.Sleep(time.Duration(5) * time.Second)

			select {

			case <-ctx.Done():
				fmt.Println("Done")

			default:
				fmt.Print("excuting")

			}
		}
	}(ctx)
	ch := <-ctx.Done()
	fmt.Print(ch)

}
