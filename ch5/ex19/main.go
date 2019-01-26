package main

import "fmt"

func main() {
	fmt.Println(one())
}

func one() (p int) {
	defer func() {
		p = recover().(int)
	}()
	panic(1)
}
