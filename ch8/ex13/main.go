package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string) // クライアントから受信する全てのメッセージ
)

func broadcaster() {
	clients := make(map[client]bool) // 全ての接続されているクライアント
	for {
		select {
		case msg := <-message:
			// 受信したメッセージを全てのクライアントの
			// 送信用メッセージチャネルへブロードキャストする
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.ch <- getNames(clients)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func getNames(clients map[client]bool) string {
	var b bytes.Buffer
	b.WriteString("members: [")
	for k := range clients {
		b.WriteString(k.name + " ")
	}
	b.WriteByte(']')
	return b.String()
}

func handleConn(conn net.Conn) {
	awake := make(chan struct{})
	go teacher(conn, awake) // 寝てないか見張っている

	ch := make(chan string) // 送信用のクライアントメッセージ
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	message <- who + " has arrived"
	entering <- client{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		awake <- struct{}{}
		message <- who + ": " + input.Text()
	}

	leaving <- client{ch, who}
	message <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func teacher(conn net.Conn, awake chan struct{}) {
	tick := time.NewTicker(10 * time.Second)
loop:
	for {
		select {
		case <-tick.C:
			break loop
		case <-awake:
			tick.Stop()
			tick = time.NewTicker(10 * time.Second)
		}
	}
	conn.Close()
}
