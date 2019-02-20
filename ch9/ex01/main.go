package main

import "fmt"

type withdraw struct {
	amount int
	ch     chan bool
}

var deposits = make(chan int) // 入金額を送信する
var balances = make(chan int) // 残高受信
var withdraws = make(chan withdraw)

func Deposits(amount int) {
	deposits <- amount
}

func Withdraw(amount int) bool {
	r := make(chan bool)
	withdraws <- withdraw{amount, r}
	return <-r
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdraws:
			if w.amount <= balance {
				balance -= w.amount
				w.ch <- true
			} else {
				w.ch <- false
			}
		}
	}
}

func main() {
	go teller()
	Deposits(100)
	for i := 0; i < 10; i++ {
		if ok := Withdraw(15); ok {
			fmt.Printf("withdraw: %v success, your balance %v\n", 15, Balance())
		} else {
			fmt.Printf("withdraw: %v fail, your balance %v\n", 15, Balance())
		}
	}
}
