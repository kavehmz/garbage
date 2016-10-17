package main

import (
	"context"
	"fmt"
	"time"
)

func t(ctx context.Context) {
	ctx2, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	select {
	case <-time.After(4 * time.Second):
		fmt.Println("overslept")
	case <-ctx2.Done():
		fmt.Println("Done with child")
		fmt.Println(ctx2.Err())
	}
	fmt.Println(ctx2.Value("a"))

	cancel()
}
func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)

	ctx = context.WithValue(ctx, "a", 42)

	go t(ctx)
	// time.Sleep(time.Millisecond * 200)
	// cancel()
	select {
	case <-time.After(4 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

	// Even though ctx should have expired already, it is good
	// practice to call its cancelation function in any case.
	// Failure to do so may keep the context and its parent alive
	// longer than necessary.
	fmt.Println(ctx.Value("a"))
	cancel()

}
