package main

import (
	"flag"
	"fmt"

	"../temperature"
)

var temp = temperature.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(temp)
}
