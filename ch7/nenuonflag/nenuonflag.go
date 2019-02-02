package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	n := NenuonFlag("nenuon", 1000000007, "?????")
	flag.Parse()
	fmt.Println(*n)
}

type Nenuon int

type nenuonFlag struct{ Nenuon }

func (n *nenuonFlag) String() string {
	return fmt.Sprintf("%d", n.Nenuon)
}

func (n *nenuonFlag) Set(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	n.Nenuon = Nenuon(i)
	return nil
}

func NenuonFlag(name string, n Nenuon, usage string) *Nenuon {
	f := nenuonFlag{n}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Nenuon
}
