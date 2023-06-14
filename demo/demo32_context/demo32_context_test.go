package demo32_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextTimeOut1(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}
func TestContextTimeOut2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func TestContext3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handle1(ctx, cancel)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func TestContext4(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handler2(ctx)
	go work(ctx, cancel)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
	time.Sleep(time.Second)
}

func TestContext5(t *testing.T) {
	var key = "key"
	ctx := context.WithValue(context.Background(), key, "hello")
	go handler3(ctx, key)
	time.Sleep(time.Second)
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func handle1(ctx context.Context, cancel func()) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	default:
		fmt.Println("default exec..")
		cancel() // 主动调用取消新号
	}
}

func handler2(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	}
}

func work(ctx context.Context, cancel func()) {
	time.Sleep(time.Second)
	cancel()
}

func handler3(ctx context.Context, key string) {
	fmt.Println(ctx.Value(key))
}
