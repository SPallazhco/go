package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "my-key", "pallazhco")
	viewContext(ctx)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go myProcess(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("We've exceeded the dead line")
		fmt.Println(ctx.Err())
	}
}

func viewContext(ctx context.Context) {
	fmt.Printf("My value is %s\n", ctx.Value("my-key"))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("PROCESS")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
