package main

import "fmt"

var deposits = make(chan int) // 入金額を送信する
var balances = make(chan int) // 残高受信
var withdrawAmount = make(chan int)
var withdrawOk = make(chan bool)

func Deposits(amount int) {
	deposits <- amount
}

func Withdraw(amount int) bool {
	withdrawAmount <- amount
	return <-withdrawOk
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
		case amount := <-withdrawAmount:
			if amount <= balance {
				balance -= amount
				withdrawOk <- true
			} else {
				withdrawOk <- false
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
