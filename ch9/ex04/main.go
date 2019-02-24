package main

import (
	"fmt"
	"time"
)

func main() {
	cnt := make(chan int)
	sum := 0
	go func() {
		for i := 0; ; i++ {
			if i%2 == 0 {
				cnt <- 1
			} else {
				<-cnt
				sum++
			}
		}
	}()

	go func() {
		for i := 0; ; i++ {
			if i%2 == 1 {
				cnt <- 1
			} else {
				<-cnt
				sum++
			}
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(sum)
}
