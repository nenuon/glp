package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"../conv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			x, _ := strconv.ParseFloat(input.Text(), 64)
			printAll(x)
		}
	}
	for _, arg := range args {
		x, _ := strconv.ParseFloat(arg, 64)
		printAll(x)
	}
}

func printAll(x float64) {
	f := conv.Fahrenheit(x)
	c := conv.Celsius(x)
	fe := conv.Feet(x)
	m := conv.Metre(x)
	p := conv.Pound(x)
	k := conv.Kilogramme(x)
	fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
	fmt.Printf("%s = %s, %s = %s\n", fe, conv.FToM(fe), m, conv.MToF(m))
	fmt.Printf("%s = %s, %s = %s\n", p, conv.PToK(p), k, conv.KToP(k))
}
