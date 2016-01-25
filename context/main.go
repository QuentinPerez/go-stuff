package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"

	"golang.org/x/net/context"
)

func printStr(ctx context.Context, wait *sync.WaitGroup, t time.Duration) {
	defer wait.Done()
	defer fmt.Println("Exit context")

	value := ctx.Value("string")
	if value == nil {
		value = "Hello World"
	}

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.Canceled:
			fmt.Println("context has been canceled")
		case context.DeadlineExceeded:
			fmt.Println("context deadline exceeded")
		default:
			logrus.Fatalf("-> %v", ctx.Err())
		}
	case <-time.After(t):
		fmt.Println(value)
	}
}

func main() {
	wait := &sync.WaitGroup{}

	fmt.Println(">>>> Context.Background")
	fmt.Println()

	wait.Add(1)
	go printStr(context.Background(), wait, 1*time.Second)
	wait.Wait()

	fmt.Println()
	fmt.Println(">>>> Context.Background with cancel function")
	fmt.Println()

	wait.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go printStr(ctx, wait, 1*time.Second)
	cancel()
	wait.Wait()

	fmt.Println()
	fmt.Println(">>>> Context.Background with cancel function + string value")
	fmt.Println()

	wait.Add(1)
	ctx, cancel = context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "string", "with value")
	go printStr(ctx, wait, 1*time.Second)
	wait.Wait()
	cancel()

	fmt.Println()
	fmt.Println(">>>> Context.Background with cancel function + string value, but canceled")
	fmt.Println()

	wait.Add(1)
	ctx, cancel = context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "string", "with value")
	go printStr(ctx, wait, 1*time.Second)
	cancel()
	wait.Wait()

	fmt.Println()
	fmt.Println(">>>> Context.Background with timeout function (1 second)")
	fmt.Println()

	wait.Add(1)
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	go printStr(ctx, wait, 10*time.Second)
	wait.Wait()
	cancel()

	fmt.Println()
	fmt.Println(">>>> Context.Background with Deadline function (1 second)")
	fmt.Println()

	wait.Add(1)
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	go printStr(ctx, wait, 10*time.Second)
	wait.Wait()
	cancel()
}
