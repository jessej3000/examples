package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx1, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	sleepAndTalk(ctx1, 10*time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
