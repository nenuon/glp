package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	data := []string{"hoge", "huga"}
	pipeline := toSingle(ctx, toParalell(ctx, toUpper(ctx, take(ctx, sendForever(ctx, data...), 20)), 4)...)
	for v := range pipeline {
		fmt.Println(v)
	}
}

// 受け取った値を送る
func send(ctx context.Context, values ...string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, v := range values {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- v
			}
		}
	}()
	return ch
}

// 受け取った値を大文字にする
func toUpper(ctx context.Context, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			select {
			case <-ctx.Done():
				return
			default:
				out <- strings.ToUpper(s)
			}
		}
	}()
	return out
}

// 受け取った値を無限に繰り返して送る
func sendForever(ctx context.Context, values ...string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				default:
					ch <- v
				}
			}
		}
	}()
	return ch
}

// n個受け取る
func take(ctx context.Context, in <-chan string, n int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
			default:
				ch <- <-in
			}
		}
	}()
	return ch
}

// n個並列で送る
func toParalell(ctx context.Context, in <-chan string, n int) []<-chan string {
	outs := make([]<-chan string, n)
	for i := 0; i < n; i++ {
		outs[i] = in
	}
	return outs
}

func toSingle(ctx context.Context, ins ...<-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		var wg sync.WaitGroup
		for i := 0; i < len(ins); i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for s := range ins[i] {
					select {
					case <-ctx.Done():
					default:
						out <- s
					}
				}
			}(i)
		}
		wg.Wait()
	}()
	return out
}
